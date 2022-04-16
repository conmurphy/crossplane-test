package application

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure ...
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("aci_application_profile", func(r *config.Resource) {
		r.ShortGroup = "application"
		r.Kind = "ApplicationProfile"
		r.References["tenant_dn"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-aci/apis/root/v1alpha1.Tenant",
		}
	})

	p.AddResourceConfigurator("aci_application_epg", func(r *config.Resource) {
		r.ShortGroup = "application"
		r.Kind = "Epg"
		r.References["application_profile_dn"] = config.Reference{
			Type: "ApplicationProfile",
		}
	})

	p.AddResourceConfigurator("aci_endpoint_security_group", func(r *config.Resource) {
		r.ShortGroup = "application"
		r.Kind = "Esg"
		r.References["application_profile_dn"] = config.Reference{
			Type: "ApplicationProfile",
		}
		r.References["relation_fv_rs_scope"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-aci/apis/network/v1alpha1.Vrf",
		}
	})

	p.AddResourceConfigurator("aci_endpoint_security_group_epg_selector", func(r *config.Resource) {
		r.ShortGroup = "application"
		r.Kind = "EsgEpgSelector"
		r.References["endpoint_security_group_dn"] = config.Reference{
			Type: "Esg",
		}
		r.References["match_epg_dn"] = config.Reference{
			Type: "Epg",
		}
	})

	p.AddResourceConfigurator("aci_endpoint_security_group_selector", func(r *config.Resource) {
		r.ShortGroup = "application"
		r.Kind = "EsgSelector"
		r.References["endpoint_security_group_dn"] = config.Reference{
			Type: "Esg",
		}
	})

	p.AddResourceConfigurator("aci_endpoint_security_group_tag_selector", func(r *config.Resource) {
		r.ShortGroup = "application"
		r.Kind = "EsgTagSelector"
		r.References["endpoint_security_group_dn"] = config.Reference{
			Type: "Esg",
		}
	})

}
