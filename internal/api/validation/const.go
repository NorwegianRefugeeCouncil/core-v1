package validation

import "github.com/nrc-no/notcore/internal/api"

var allowedDisplacementStatuses = api.AllDisplacementStatuses()
var allowedGenders = api.AllGenders()

var allowedDisplacementStatusesStr []string
var allowedGendersStr []string

func init() {
	for _, v := range allowedDisplacementStatuses.Items() {
		allowedDisplacementStatusesStr = append(allowedDisplacementStatusesStr, string(v))
	}
	for _, v := range allowedGenders.Items() {
		allowedGendersStr = append(allowedGendersStr, string(v))
	}
}
