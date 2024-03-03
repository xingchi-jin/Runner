package internalvars

import (
	"errors"
	"fmt"
)

type Var struct {
	name         string
	exported     bool
	value        string
	confidential bool
}
type Variables map[string]*Var

type GroupVars struct {
	groupVars map[string]Variables
}

func NewGroupVars() *GroupVars {
	return &GroupVars{
		groupVars: make(map[string]Variables),
	}
}

func (ev *GroupVars) Add(task_name, var_name, val string, exported, confidential bool) {
	task_vars, ok := ev.groupVars[task_name]
	if ok == false {
		ev.groupVars[task_name] = make(map[string]*Var)
	}
	task_vars[var_name] = &Var{
		name:     var_name,
		exported: exported,
		value:    val,
	}
}

func (ev *GroupVars) Get(task_name, var_name string) (string, error) {
	if task_vars, ok := ev.groupVars[task_name]; ok {
		if variable, var_ok := task_vars[var_name]; var_ok {
			return variable.value, nil
		}
	}
	return "", errors.New(fmt.Sprintf("%s:%s Not Found", task_name, var_name))
}
