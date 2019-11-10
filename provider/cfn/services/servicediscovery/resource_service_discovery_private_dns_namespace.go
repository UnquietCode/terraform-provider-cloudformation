// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package servicediscovery

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceDiscoveryPrivateDnsNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDiscoveryPrivateDnsNamespaceCreate,
		Read:   resourceServiceDiscoveryPrivateDnsNamespaceRead,
		Update: resourceServiceDiscoveryPrivateDnsNamespaceUpdate,
		Delete: resourceServiceDiscoveryPrivateDnsNamespaceDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"vpc": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceDiscoveryPrivateDnsNamespaceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceDiscovery::PrivateDnsNamespace", data, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceDiscovery::PrivateDnsNamespace", data, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceDiscovery::PrivateDnsNamespace", data, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceDiscovery::PrivateDnsNamespace", data, meta)
}