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
	val, err := im.groupVars.Get(task_name, variable_name)
	if err != nil {
		return err
	}
	_, ok := im.importedVars[task_name]
	if ok == false {
		im.importedVars[task_name] = make(map[string]string)
	}
	im.importedVars[task_name][variable_name] = val
	return nil
}

func (im *Imports) GetAll() map[string]map[string]string {
	return im.importedVars
}
