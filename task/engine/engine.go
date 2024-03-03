package engine

import (
	"context"
	"fmt"
	"net/http"

	"github.com/harness/runner/router"
	"github.com/harness/runner/task"
	"github.com/harness/runner/task/helper"
	"github.com/harness/runner/task/internalvars"
	"github.com/sirupsen/logrus"
)

func Process(taskGroup *task.TaskGroup, router router.Router) (*task.TaskGroupResult, error) {
	// TODO: Run graph processing to reorder tasks sequence
	// graph.process(request.Tasks)
	tasks := taskGroup.Tasks
	var response http.ResponseWriter
	var err error
	ctx := context.Background()
	group_vars := internalvars.NewGroupVars()

	// TODO: execute independent subsets in parallel. The for loop below processes one subset
	for _, this_task := range tasks {
		// Retrieve import requested variables
		imports := internalvars.NewImportedVars(group_vars)
		for _, api_imports := range this_task.Imports {
			err := imports.Add(api_imports.TaskName, api_imports.VariableName)
			if err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{
					"taskName":  this_task.Name,
					"taskType":  this_task.TaskType,
					"importVar": fmt.Sprintf("%s.%s", api_imports.TaskName, api_imports.VariableName),
				}).Error("Resolve import failed")
				// TODO: fill task group result
				return nil, err
			}
		}
		if response, err = processTask(ctx, this_task, imports, router); err != nil {
			logrus.WithError(err).Error("Executing task failed")
			// TODO: fill task group result
			return nil, err
		} else {
			// Read outputs from response and populate the group vars
			env_outputs, err = response.ReadBodyAsMap()
			if err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{
					"taskName": this_task.Name,
					"taskType": this_task.TaskType,
				}).Error("Read task output failed")
				// TODO: fill task group result
				return nil, err
			}
			// Read exports
			tmp_exported_vars := make(map[string]*task.Export)
			for _, api_export := range this_task.Exports {
				tmp_exported_vars[api_export.VariableName] = api_export
			}
			// Add output vars to the group variable list.
			for output_var_name, output_var_value := range env_outputs {
				var is_exported, is_confidential bool = false, false
				if ex, ok := tmp_exported_vars[output_var_name]; ok {
					is_exported = true
					is_confidential = ex.Confidential
				}
				group_vars.Add(this_task.Name, output_var_name, output_var_value, is_exported, is_confidential)
			}
		}
	}
	// TODO: fill task group result
	return nil, nil
}

func processTask(ctx context.Context, task *task.Task, imports *internalvars.Imports, router router.Router) (*helper.Response, error) {
	// Based on task type, invoke corresponding handler
	req, err := helper.NewRequest(ctx, task.Spec, imports)
	if err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"taskId":   task.Name,
			"taskName": task.Name,
		}).Error("Create internal request object failed")
		return nil, err
	}
	resp := helper.NewResponseWriter()
	// Invoke http handler locally, or (TODO) invoke remote runner api
	router.Route(task.TaskType).ServeHTTP(resp, req)

	return resp, nil
}
