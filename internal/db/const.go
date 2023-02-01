package db

import "time"

// maxParams is the maximum number of arguments that can be passed to a postgres query
const maxParams = 65535

type individualAction struct {
	conditions  []string
	targetField string
	newValue    any
}

var deleteAction = individualAction{
	conditions:  []string{"and deleted_at IS NULL"},
	targetField: "deleted_at",
	newValue:    time.Now().UTC().Format(time.RFC3339),
}
var activateAction = individualAction{
	conditions:  []string{},
	targetField: "inactive",
	newValue:    false,
}
var deactivateAction = individualAction{
	conditions:  []string{},
	targetField: "inactive",
	newValue:    true,
}

const (
	DeleteAction     string = "delete"
	ActivateAction          = "activate"
	DeactivateAction        = "deactivate"
)

var individualActions = map[string]individualAction{
	DeleteAction:     deleteAction,
	ActivateAction:   activateAction,
	DeactivateAction: deactivateAction,
}
