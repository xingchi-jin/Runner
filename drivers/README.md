# Drivers
Driver interface allows a user to implement a customized task type. Driver serves as an infrastructure manager that can be used to manage infra and process tasks executed on those infra.

Driver design follows similar pattern as Nomad driver.

## Usage
Here in this package, it provides a super-simplified interface design, to explain the idea.

Step 1. A user can implement a new driver by implementing the DriverPlugin interface.

```
type HelloWorldDriverPlugin struct {
}

func (dp *HelloWorldDriverPlugin) GetDriverConfig() *Driver {
	return & Driver {
		task_type: "helloworld_driver",
	}
}

func (dp *HelloWorldDriverPlugin) ExecTask(input []byte) (*ExecTaskResult, error) {
	// implementation detail
	fmt.Println("helloworld")

	return &ExecTaskResult {
		results: map[string]string {
			"output1": "1",
			"output2": "2",
		}, 
	}, nil
}

func (dp *HelloWorldDriverPlugin) HanHandleSignals(sigChan <- chan Signal) {
    switch sig := <- sigChan; sig {
        case "destroy":
          ...
        case "start":
          ...
        case "pause":
          ...
    }
}
```

Step 2. After the driver is registered, we can use the driver to execute tasks. Example as below

```
taskGroup: 
- name: foo,
  type: helloworld_driver,
  spec:
    // any data
  exports:
  - name: output1
- name: bar,
  type: os_exec,
  spec:
    command: ["/bin/bash" "-c", "echo '{{ .foo.output1 )}'"],
    envs:
    - key: value
  imports:
  - from: foo,
    variable: output1
```