package root

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure ...
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aci_tenant", func(r *config.Resource) {
		r.ShortGroup = "root"
		r.Kind = "Tenant"
	})

}
