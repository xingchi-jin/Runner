# Harness Expression Evaluation Task
Expression evaluation task interpolates expression from an input string. It provides evaluated string in the output.

The expression is a Harness compatible expression that can be used to refer to Harness input, output, and execution variables. Please read [Harness Expression](https://developer.harness.io/docs/platform/variables-and-expressions/runtime-inputs/#expressions)

This task provides the ability to do expression evaluation in the Runner. Use cases like, the outcome from a previous Harness task has expression in it, while the next Harness task requires this outcome as an input. In this case, the expression evaluation can happen in the middle of executing these two tasks. This avoids the back-and-forth communications to send the intermediate outcome back to server to do evaluation, then pass back to the same Runner for executing the next task.

## Why Harness Expression Evaluation is a task
Expression evaluation for harness variables is better to be a task type instead of an external plugin. Because it requires dial home ability to connect to Harness server, while we expect plugin to be agnostic of Harness's authentication mechanisms. Implementing it as a task allow the evaluation to happen on top of an already-established secure communication channel.
Plugins no longer needs to worry about secure network connections, no Harness credentials will be leaked to plugins.

## Sample
```
taskGroup:
- name: fetch-k8s-manifest
  type: os_exec
  spec:
    command: |-
      curl https://...k8s-deployment.yaml | echo "yaml_data: $@-" > $OUTPUT_FILE
  exports:
    - name: yaml_data
- name: resolve,
  type: expression_evalation,
  spec:
    string_to_evaluate: {{ .fetch-k8s-manifest.yaml_data }}
  imports:
    - from: fetch-k8s-manifest
      name: yaml_data
      scope: local
  exports:
    - name: result
      scope: local
- name: deploy
  type: os_exec
  spec:
    command: |-
      cat < EOF | kubectl apply -f -
      {{ .resolve.result }}
      EOF
  imports:
    - task_name: resolve,
      name: result
```
This example explains a scenario where the client wants to download a k8s manifest yaml while there are expressions in it. This above tasks definition will download the file --> resolve expression --> deploy
