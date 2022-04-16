package network

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure ...
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aci_vrf", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.Kind = "Vrf"
	})
	p.AddResourceConfigurator("aci_subnet", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.Kind = "Subnet"
		r.References["parent_dn"] = config.Reference{
			Type: "Vrf",
		}
	})
	p.AddResourceConfigurator("aci_bridge_domain", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.Kind = "BridgeDomain"
		r.References["tenant_dn"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-aci/apis/root/v1alpha1.Tenant",
		}
	})
	p.AddResourceConfigurator("aci_l3_outside", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.Kind = "L3Outside"
		r.References["tenant_dn"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-aci/apis/root/v1alpha1.Tenant",
		}
	})
	p.AddResourceConfigurator("aci_external_network_instance_profile", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.Kind = "L3ExternalNetworkInstanceProfile"
		r.References["l3_outside_dn"] = config.Reference{
			Type: "L3Outside",
		}
	})
	p.AddResourceConfigurator("aci_l3_ext_subnet", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.Kind = "L3ExternalSubnet"
		r.References["external_network_instance_profile_dn"] = config.Reference{
			Type: "L3ExternalNetworkInstanceProfile",
		}
	})
	p.AddResourceConfigurator("aci_logical_node_profile", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.Kind = "L3LogicalNodeProfile"
		r.References["l3_outside_dn"] = config.Reference{
			Type: "L3Outside",
		}
	})
	p.AddResourceConfigurator("aci_logical_interface_profile", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.Kind = "L3LogicalInterfaceProfile"
		r.References["logical_node_profile_dn"] = config.Reference{
			Type: "L3LogicalNodeProfile",
		}
	})

}
