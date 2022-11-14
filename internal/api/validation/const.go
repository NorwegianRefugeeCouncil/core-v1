package validation

import "github.com/nrc-no/notcore/internal/api"

var allowedDisplacementStatuses = api.AllDisplacementStatuses()
var allowedSexs = api.AllSexs()

var allowedDisplacementStatusesStr []string
var allowedSexsStr []string

var allowedDisabilityLevels = api.AllDisabilityLevels()
var allowedDisabilityLevelsStr []string

func init() {
	for _, v := range allowedDisplacementStatuses.Items() {
		allowedDisplacementStatusesStr = append(allowedDisplacementStatusesStr, string(v))
	}
	for _, v := range allowedSexs.Items() {
		allowedSexsStr = append(allowedSexsStr, string(v))
	}
	for _, v := range allowedDisabilityLevels.Items() {
		allowedDisabilityLevelsStr = append(allowedDisabilityLevelsStr, string(v))
	}
}
