package expreval

var TaskName string = "expression_evalation"

type ExpressionEvaluationTaskHandler struct {
}

type Spec struct {
	Context          *HarnessContext   `json:"harness_context"`
	StringToEvaluate string            `json:"string_to_evaluate"`
	Env              map[string]string `json:"data"`
}

type HarnessContext struct {
	AccountId string `json:"account_id"`
	OrgId     string `json:"organization_id"`
	ProjectId string `json:"project_id"`
}
