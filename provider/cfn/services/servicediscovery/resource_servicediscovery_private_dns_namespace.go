// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-privatednsnamespace.html

package servicediscovery

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceDiscoveryPrivateDnsNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDiscoveryPrivateDnsNamespaceCreate,
		Read:   resourceServiceDiscoveryPrivateDnsNamespaceRead,
		Delete: resourceServiceDiscoveryPrivateDnsNamespaceDelete,

		Schema: map[string]*schema.Schema{
			"the_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
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
	return plugin.ResourceCreate("AWS::ServiceDiscovery::PrivateDnsNamespace", ResourceServiceDiscoveryPrivateDnsNamespace(), data, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceDiscovery::PrivateDnsNamespace", ResourceServiceDiscoveryPrivateDnsNamespace(), data, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceDiscovery::PrivateDnsNamespace", ResourceServiceDiscoveryPrivateDnsNamespace(), data, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceDiscovery::PrivateDnsNamespace", data, meta)
}