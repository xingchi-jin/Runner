package vault

type HashicorpVaultHandler struct{}

var TaskName string = "os_exec"

type Spec struct {
	Url        string    `json:"url"`
	Role       string    `json:"role"`
	Path       string    `json:"path"`
	SecretName string    `json:"secret_name"`
	Artifact   *Artifact `json:",omitempty"`
}

type Artifact struct {
	Source string `json:"source"` // container or binary
}
