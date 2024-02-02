package db

import "time"

// maxParams is the maximum number of arguments that can be passed to a postgres query
const maxParams = 65535

type individualActionConfig struct {
	conditions  []string
	targetField string
	newValue    any
}

var deleteAction = individualActionConfig{
	conditions:  []string{"and deleted_at IS NULL"},
	targetField: "deleted_at",
	newValue:    time.Now().UTC().Format(time.RFC3339),
}
var activateAction = individualActionConfig{
	conditions:  []string{},
	targetField: "inactive",
	newValue:    false,
}
var deactivateAction = individualActionConfig{
	conditions:  []string{},
	targetField: "inactive",
	newValue:    true,
}

type IndividualAction string

const (
	DeleteAction     IndividualAction = "delete"
	ActivateAction   IndividualAction = "activate"
	DeactivateAction IndividualAction = "deactivate"
)

var individualActionsConfig = map[IndividualAction]individualActionConfig{
	DeleteAction:     deleteAction,
	ActivateAction:   activateAction,
	DeactivateAction: deactivateAction,
}
