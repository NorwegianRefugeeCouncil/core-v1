package db

// maxParams is the maximum number of arguments that can be passed to a postgres query
const maxParams = 65535

type IndividualAction string

const (
	DeleteAction     IndividualAction = "delete"
	ActivateAction   IndividualAction = "activate"
	DeactivateAction IndividualAction = "deactivate"
)
