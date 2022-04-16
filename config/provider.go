/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-jet-aci/config/application"
	"github.com/crossplane-contrib/provider-jet-aci/config/network"
	"github.com/crossplane-contrib/provider-jet-aci/config/root"
)

const (
	resourcePrefix = "aci"
	modulePath     = "github.com/crossplane-contrib/provider-jet-aci"
)

//go:embed schema.json
var providerSchema string

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithIncludeList([]string{
			// "aci_aaa_domain$",
			// "aci_aaa_domain_relationship$",
			// "aci_access_generic$",
			// "aci_access_group$",
			// "aci_access_port_block$",
			// "aci_access_port_selector$",
			// "aci_access_sub_port_block$",
			// "aci_access_switch_policy_group$",
			// "aci_action_rule_profile$",
			// "aci_annotation$",
			// "aci_any$",
			"aci_application_epg$",
			"aci_application_profile$",
			// "aci_attachable_access_entity_profile$",
			// "aci_authentication_properties$",
			// "aci_autonomous_system_profile$",
			// "aci_bd_dhcp_label$",
			// "aci_bfd_interface_policy$",
			// "aci_bgp_address_family_context$",
			// "aci_bgp_best_path_policy$",
			// "aci_bgp_peer_connectivity_profile$",
			// "aci_bgp_peer_prefix$",
			// "aci_bgp_route_control_profile$",
			// "aci_bgp_route_summarization$",
			// "aci_bgp_timers$",
			"aci_bridge_domain$",
			// "aci_cdp_interface_policy$",
			// "aci_client_end_point$",
			// "aci_cloud_applicationcontainer$",
			// "aci_cloud_availability_zone$",
			// "aci_cloud_aws_provider$",
			// "aci_cloud_cidr_pool$",
			// "aci_cloud_context_profile$",
			// "aci_cloud_domain_profile$",
			// "aci_cloud_endpoint_selector$",
			// "aci_cloud_endpoint_selectorfor_external_epgs$",
			// "aci_cloud_epg$",
			// "aci_cloud_external_epg$",
			// "aci_cloud_provider_profile$",
			// "aci_cloud_providers_region$",
			// "aci_cloud_subnet$",
			// "aci_cloud_vpn_gateway$",
			// "aci_configuration_export_policy$",
			// "aci_configuration_import_policy$",
			// "aci_connection$",
			// "aci_console_authentication$",
			// "aci_contract$",
			// "aci_contract_subject$",
			// "aci_coop_policy$",
			// "aci_default_authentication$",
			// "aci_destination_of_redirected_traffic$",
			// "aci_dhcp_option$",
			// "aci_dhcp_option_policy$",
			// "aci_dhcp_relay_policy$",
			// "aci_duo_provider_group$",
			// "aci_encryption_key$",
			// "aci_end_point_retention_policy$",
			// "aci_endpoint_controls$",
			// "aci_endpoint_ip_aging_profile$",
			// "aci_endpoint_loop_protection$",
			"aci_endpoint_security_group$",
			"aci_endpoint_security_group_epg_selector$",
			"aci_endpoint_security_group_selector$",
			"aci_endpoint_security_group_tag_selector$",
			// "aci_epg_to_contract$",
			// "aci_epg_to_domain$",
			// "aci_epg_to_static_path$",
			// "aci_epgs_using_function$",
			// "aci_error_disable_recovery$",
			"aci_external_network_instance_profile$",
			// "aci_fabric_if_pol$",
			// "aci_fabric_node$",
			// "aci_fabric_node_control$",
			// "aci_fabric_node_member$",
			// "aci_fabric_path_ep$",
			// "aci_fabric_wide_settings$",
			// "aci_fc_domain$",
			// "aci_fex_bundle_group$",
			// "aci_fex_profile$",
			// "aci_file_remote_path$",
			// "aci_filter$",
			// "aci_filter_entry$",
			// "aci_firmware_download_task$",
			// "aci_firmware_group$",
			// "aci_firmware_policy$",
			// "aci_function_node$",
			// "aci_global_security$",
			// "aci_hsrp_group_policy$",
			// "aci_hsrp_interface_policy$",
			// "aci_imported_contract$",
			// "aci_interface_blacklist$",
			// "aci_interface_fc_policy$",
			// "aci_isis_domain_policy$",
			// "aci_l2_domain$",
			// "aci_l2_interface_policy$",
			// "aci_l2_outside$",
			// "aci_l2out_extepg$",
			// "aci_l3_domain_profile$",
			"aci_l3_ext_subnet$",
			// "aci_l3_interface_policy$",
			"aci_l3_outside$",
			// "aci_l3out_bfd_interface_profile$",
			// "aci_l3out_bgp_external_policy$",
			// "aci_l3out_bgp_protocol_profile$",
			// "aci_l3out_floating_svi$",
			// "aci_l3out_hsrp_interface_group$",
			// "aci_l3out_hsrp_interface_profile$",
			// "aci_l3out_hsrp_secondary_vip$",
			// "aci_l3out_loopback_interface_profile$",
			// "aci_l3out_ospf_external_policy$",
			// "aci_l3out_ospf_interface_profile$",
			// "aci_l3out_path_attachment$",
			// "aci_l3out_path_attachment_secondary_ip$",
			// "aci_l3out_route_tag_policy$",
			// "aci_l3out_static_route$",
			// "aci_l3out_static_route_next_hop$",
			// "aci_l3out_vpc_member$",
			// "aci_l4_l7_service_graph_template$",
			// "aci_lacp_policy$",
			// "aci_ldap_group_map$",
			// "aci_ldap_group_map_rule$",
			// "aci_ldap_group_map_rule_to_group_map$",
			// "aci_ldap_provider$",
			// "aci_leaf_access_bundle_policy_group$",
			// "aci_leaf_access_port_policy_group$",
			// "aci_leaf_breakout_port_group$",
			// "aci_leaf_interface_profile$",
			// "aci_leaf_profile$",
			// "aci_leaf_selector$",
			// "aci_lldp_interface_policy$",
			// "aci_local_user$",
			// "aci_logical_device_context$",
			// "aci_logical_interface_context$",
			"aci_logical_interface_profile$",
			"aci_logical_node_profile$",
			// "aci_logical_node_to_fabric_node$",
			// "aci_login_domain$",
			// "aci_login_domain_provider$",
			// "aci_maintenance_group_node$",
			// "aci_maintenance_policy$",
			// "aci_managed_node_connectivity_group$",
			// "aci_match_route_destination_rule$",
			// "aci_match_rule$",
			// "aci_mcp_instance_policy$",
			// "aci_mgmt_preference$",
			// "aci_mgmt_zone$",
			// "aci_miscabling_protocol_interface_policy$",
			// "aci_monitoring_policy$",
			// "aci_node_block$",
			// "aci_node_block_firmware$",
			// "aci_node_mgmt_epg$",
			// "aci_ospf_interface_policy$",
			// "aci_ospf_route_summarization$",
			// "aci_ospf_timers$",
			// "aci_physical_domain$",
			// "aci_pod_maintenance_group$",
			// "aci_port_security_policy$",
			// "aci_port_tracking$",
			// "aci_qos_instance_policy$",
			// "aci_radius_provider$",
			// "aci_radius_provider_group$",
			// "aci_ranges$",
			// "aci_recurring_window$",
			// "aci_rest$",
			// "aci_rest_managed$",
			// "aci_route_control_context$",
			// "aci_route_control_profile$",
			// "aci_rsa_provider$",
			// "aci_saml_certificate$",
			// "aci_saml_provider$",
			// "aci_saml_provider_group$",
			// "aci_service_redirect_policy$",
			// "aci_snmp_community$",
			// "aci_span_destination_group$",
			// "aci_span_source_group$",
			// "aci_span_sourcedestination_group_match_label$",
			// "aci_spanning_tree_interface_policy$",
			// "aci_spine_access_port_selector$",
			// "aci_spine_interface_profile$",
			// "aci_spine_interface_profile_selector$",
			// "aci_spine_port_policy_group$",
			// "aci_spine_port_selector$",
			// "aci_spine_profile$",
			// "aci_spine_switch_association$",
			// "aci_spine_switch_policy_group$",
			// "aci_static_node_mgmt_address$",
			// "aci_subnet$",
			// "aci_system$",
			// "aci_taboo_contract$",
			// "aci_tacacs_accounting$",
			// "aci_tacacs_accounting_destination$",
			// "aci_tacacs_provider$",
			// "aci_tacacs_provider_group$",
			// "aci_tacacs_source$",
			// "aci_tag$",
			"aci_tenant$",
			// "aci_trigger_scheduler$",
			// "aci_user_security_domain$",
			// "aci_user_security_domain_role$",
			// "aci_vlan_encapsulationfor_vxlan_traffic$",
			// "aci_vlan_pool$",
			// //"aci_vmm_controller$",
			// "aci_vmm_credential$",
			// "aci_vmm_domain$",
			// "aci_vpc_domain_policy$",
			// "aci_vpc_explicit_protection_group$",
			"aci_vrf$",
			// "aci_vrf_snmp_context$",
			// "aci_vrf_snmp_context_community$",
			// "aci_vrf_to_bgp_address_family_context$",
			// "aci_vsan_pool$",
			// "aci_vswitch_policy$",
			// "aci_vxlan_pool$",
			// "aci_x509_certificate$",
		}),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		root.Configure,
		application.Configure,
		network.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
