package convert

import (
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/storage"
)

// HalToKeel generates the Spinnaker keel config for the supplied config.Hal h
func HalToKeel(h *config.Hal) *config.Keel {
	if !h.GetManagedDelivery().GetEnabled().GetValue() {
		return &config.Keel{}
	}

	var sql *config.Keel_SQL
	if h.GetPersistentStorage().GetSql().GetEnabled().GetValue() {
		sql = &config.Keel_SQL{
			Enabled: h.GetPersistentStorage().GetSql().GetEnabled(),
			ConnectionPools: &storage.ConnectionPools{
				Default: h.GetPersistentStorage().GetSql().GetDefault(),
			},
			Migration: h.GetPersistentStorage().GetSql().GetMigration(),
		}
	}

	return &config.Keel{
		Sql:    sql,
		Keel:   h.GetManagedDelivery().GetKeel(),
		Eureka: h.GetManagedDelivery().GetEureka(),
	}
}
