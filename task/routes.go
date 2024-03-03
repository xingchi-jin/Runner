package task

import (
	"github.com/harness/runner/taskimpl/osexec"
)

var routes map[string]Handler

func GetTasksHandlers() map[string]Handler {
	if routes == nil {
		routes = make(map[string]Handler)
		routes["os_exec"] = &osexec.OsExecHandler{}
	}
	return routes
}
