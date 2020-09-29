package convert

import (
	"github.com/spinnaker/kleat/api/client/config"
)

// HalToKeel generates the Spinnaker keel config for the supplied config.Hal h
func HalToKeel(h *config.Hal) *config.Keel {
	if !h.GetManagedDelivery().GetEnabled().GetValue() {
		return &config.Keel{}
	}

	return &config.Keel{
		Sql:    h.GetPersistentStorage().GetSql(),
		Keel:   h.GetManagedDelivery().GetKeel(),
		Eureka: h.GetManagedDelivery().GetEureka(),
	}
}
