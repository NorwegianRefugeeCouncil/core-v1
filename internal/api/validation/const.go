package validation

import (
	"github.com/nrc-no/notcore/internal/api/enumTypes"
)

var allowedDisplacementStatuses = enumTypes.AllDisplacementStatuses()
var allowedSexes = enumTypes.AllSexes()

var allowedDisplacementStatusesStr []string
var allowedSexesStr []string

var allowedIdentificationTypes = enumTypes.AllIdentificationTypes()
var allowedIdentificationTypesStr []string

var allowedContactMethods = enumTypes.AllContactMethods()
var allowedContactMethodsStr []string

var allowedDisabilityLevels = enumTypes.AllDisabilityLevels()
var allowedDisabilityLevelsStr []string

var allowedServiceCCs = enumTypes.AllServiceCCs()
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
