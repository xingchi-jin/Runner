package run

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/harness/runner/router"
	"github.com/harness/runner/task"
	"github.com/harness/runner/task/engine"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

type RunCommand struct {
	taskFile string
}

func (c *RunCommand) run(*kingpin.ParseContext) error {
	// Parse tasks definition file
	data, err := os.ReadFile(c.taskFile)
	if err != nil {
		logrus.WithError(err).Error("Read task file failed")
		return err
	}

	request := &task.TaskGroup{}
	if err := json.Unmarshal(data, request); err != nil {
		logrus.WithError(err).Error("Unmarshal task file to json failed")
		return err
	}

	router := router.NewRouter(task.GetTasksHandlers())

	// It makes sense for one DAG to share the same context
	response, _ := engine.Process(request, router)
	fmt.Print(response)
	return nil
	// var response http.ResponseWriter

	// for _, task := range tasks {
	// 	if response != nil {
	// 		// write response of previous task into the current one's request
	// 	}
	// 	if response, err = processTask(ctx, task, imports, router); err != nil {
	// 		logrus.WithError(err).Error("Executing task failed")
	// 		return err
	// 	}
	// }

	// // Read output from cgi response

}

// Register the server commands.
func Register(app *kingpin.Application) {
	c := new(RunCommand)

	cmd := app.Command("run", "Run tasks").Action(c.run)
	cmd.Flag("tasks", "tasks definition file").StringVar(&c.taskFile)
}
