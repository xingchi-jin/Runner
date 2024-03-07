# Hashicorp Vault Task
Hashicorp Vault task provides the ability to fetch secret from Hashicorp Vault. Besides of a default implementation to fetch secrets, it supports invoking an external executable to fetch secrets

## Sample
```
taskGroup:
- name: fetch-secret
  type: vault
  spec:
    secret_name: password
    path: ...
    role: ...
    url: ...
    artifact:
      source: https://...custom.jar
  exports:
    - name: password
      scope: local
      confidential: true  // it will be masked in logs
- name: curl_with_password,
  type: os_exec,
  spec:
    command: ["/bin/bash" "-c", "echo '{{ .fetch-secret.password )}'"],
    envs:
    - key: value
  imports:
  - from: foo,
    variable: password
```
This example explains a use case where a curl command execution requires some password.
1. This password needs to come from a vault.
2. In order fetch the secret, the client prefers to use custom.jar instead of using a default implementation. 
3. All the inputs like `secret_name` `path` `role` will be provided to the custom.jar for execution, how it's provided is up to the implementation details of the task `vault`
