# Cache
Runner's cache temporarily stores the outputs of a task. A task execution will be skipped if the outputs of the task is already present in the cache.

To support this, we introduce a `cache` declarative in the [Task](../README.md#tasks) definition.

## Sample
```
taskGroup: 
- name: foo,
  type: os_exec,
  cache:
    key: "mycorp.vault-1"
  spec:
    command: |
      secret=curl https://vault...;
      echo "password=${secret}" > ${OUTPUT_FILE}_foo;
  exports:
  - name: password
    confidential: true
```
In this example, when Runner detects the cache key "mycorp.vault-1" present in the cache, it will skip the execution of "foo". The outputs of this task will be retrieved from the cache.