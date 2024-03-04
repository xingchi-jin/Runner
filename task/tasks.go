package task

type Task struct {
	Name      string      `json:"name"` // Name be unique within a group
	Id        string      `json:"id,omitempty"`
	TaskType  string      `json:"type"`
	Spec      interface{} `json:"spec"` // Task's payload
	DependsOn []string    `json:"depends_on"`
	Exports   []*Export   `json:"exports"`
	Imports   []*Import   `json:"imports"`
	// TODO: ForwardTo *Forward `json:"forward_to"`
}

// type Forward struct {
// 	Host string `json:"host"`
// 	// TODO: credentials ?
// }

type Export struct {
	VariableName string `json:"name"`
	Confidential bool   `json:"confidential"`
}

type Import struct {
	TaskName     string `json:"task_name"`
	VariableName string `json:"name"`
}

type TaskGroup struct {
	Tasks []*Task `json:"tasks"`
	// TODO: ForwardTo *Forward `json:"forward_to"`
}

type TaskStatus string

const (
	TASK_STATUS_SUCCESS TaskStatus = "success"
	TASK_STATUS_FAILED  TaskStatus = "failed"
	TASK_STATUS_SKIPPED TaskStatus = "skipped"
)

type taskResult struct {
	TaskName string            `json:"name"`
	Id       string            `json:"id,omitempty"`
	Status   TaskStatus        `json:"status"`
	Output   map[string]string `json:"output"`
}

type TaskGroupResult struct {
	Results map[string]*taskResult `json:"results"`
}
