package driver

type RunnerClient interface {
	RegisterDriver(driver *Driver)
}

type Signal string

const (
	start Signal = "start"
	pause Signal = "pause"
)

type Driver struct {
	task_type string `json:"task_type"`
}

type ExecTaskResult struct {
	results map[string]string
}

type DriverPlugin interface {
	GetDriverConfig() *Driver

	ExecTask(input []byte) (*ExecTaskResult, error)

	HandleSignals(sigChan <-chan Signal)
}
