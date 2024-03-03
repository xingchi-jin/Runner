# Runner
- [Runner](#runner)
  - [About](#about)
  - [Tasks](#tasks)
      - [Schema](#schema)
    - [Task Types](#task-types)
    - [Task Group](#task-group)
      - [Schema](#schema-1)
    - [Sample](#sample)
## About
Runner support executing tasks in various environments, for example k8s, docker, MacOs, Linux and Windows.

## Tasks
Runner is a general task engine, it executes tasks. Runner natively implements a list of task types.

#### Schema
| config | description
| ---  | ---
| name | Task name. Task name should be unique within a task group
| type | Task type
| spec | Inputs of the task
| imports | List of variables to import
| imports.[*].from | The task name where the variable is from
| imports.[*].variable | The name of the variable to import
| exports | List of output variables to export. Once exported, those variables will be visible to subsequent tasks 
| exports.[*].name | The name of the variable to export
| exports.[*].confidential | if `true`, the variable will be masked in logs 
| depends_on | Specifies list of tasks need to execute before this one
| forward | Specifies where to execute the task.
| forward.to | IP or Hostname of the Runner where the task will be relayed to for execution.

### Task Types
"os_exec" task provides the ability to run os commands.

### Task Group
A task group consists of multiple tasks. Tasks in the task group can have causal dependencies. For example, the input of one task may require the output variables from a previous task. The other example is, one task requires the downloaded artifact from a previous task.

To support variable passing, we provide `export` and `import` declarative. See [example](###Sample).

To support artifact sharing, we provide `depends_on` declarative. This will make sure the dependent tasks will be executed beforehand. See [example](###Sample).

#### Schema
| config | description
| --- | ---
| taskGroup | A list of tasks to execute
| forward | Specifies where to execute the task.
| forward.to | IP or Hostname of the Runner where the task will be relayed to for execution.

### Sample 
```
taskGroup: 
- name: foo,
  type: os_exec,
  spec:
    command: |
      secret=curl https://vault...;
      echo "password=${secret}" > ${OUTPUT_FILE}_foo;
  exports:
  - name: password
    confidential: true
- name: bar,
  type: os_exec,
  spec:
    command: ["/bin/bash" "-c", "echo '{{ .foo.output.password )}'"],
    envs:
    - key: value  
  depends_on: foo,
  imports:
  - from: foo,
    variable: password
```
