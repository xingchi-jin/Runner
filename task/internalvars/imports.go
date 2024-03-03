package internalvars

type Imports struct {
	importedVars map[string]map[string]string
	groupVars    *GroupVars
}

func NewImportedVars(groupVars *GroupVars) *Imports {
	return &Imports{
		importedVars: make(map[string]map[string]string),
		groupVars:    groupVars,
	}
}

func (im *Imports) Add(task_name, variable_name string) error {
	return nil
}

func (im *Imports) GetAll() map[string]map[string]string {
	return nil
}
