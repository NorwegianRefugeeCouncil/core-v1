package validation

import "github.com/nrc-no/notcore/internal/api"

var allowedDisplacementStatuses = api.AllDisplacementStatuses()
var allowedSexes = api.AllSexes()

var allowedDisplacementStatusesStr []string
var allowedSexesStr []string

var allowedDisabilityLevels = api.AllDisabilityLevels()
var allowedDisabilityLevelsStr []string

var allowedServiceCCs = api.AllServiceCCs()
var allowedServiceCCsStr []string

func init() {
	for _, v := range allowedDisplacementStatuses.Items() {
		allowedDisplacementStatusesStr = append(allowedDisplacementStatusesStr, string(v))
	}
	for _, v := range allowedSexes.Items() {
		allowedSexesStr = append(allowedSexesStr, string(v))
	}
	for _, v := range allowedDisabilityLevels.Items() {
		allowedDisabilityLevelsStr = append(allowedDisabilityLevelsStr, string(v))
	}
	for _, v := range allowedServiceCCs.Items() {
		allowedServiceCCsStr = append(allowedServiceCCsStr, string(v))
	}
}
