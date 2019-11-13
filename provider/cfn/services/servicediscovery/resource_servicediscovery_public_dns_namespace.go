// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-publicdnsnamespace.html

package servicediscovery

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceDiscoveryPublicDnsNamespace() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceDiscoveryPublicDnsNamespaceExists,
		Read:   resourceServiceDiscoveryPublicDnsNamespaceRead,
		Create: resourceServiceDiscoveryPublicDnsNamespaceCreate,
		Update: resourceServiceDiscoveryPublicDnsNamespaceUpdate,
		Delete: resourceServiceDiscoveryPublicDnsNamespaceDelete,
		CustomizeDiff: resourceServiceDiscoveryPublicDnsNamespaceCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceDiscoveryPublicDnsNamespaceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceDiscoveryPublicDnsNamespaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceDiscovery::PublicDnsNamespace", ResourceServiceDiscoveryPublicDnsNamespace(), data, meta)
}

func resourceServiceDiscoveryPublicDnsNamespaceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceDiscovery::PublicDnsNamespace", ResourceServiceDiscoveryPublicDnsNamespace(), data, meta)
}

func resourceServiceDiscoveryPublicDnsNamespaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceDiscovery::PublicDnsNamespace", ResourceServiceDiscoveryPublicDnsNamespace(), data, meta)
}

func resourceServiceDiscoveryPublicDnsNamespaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceDiscovery::PublicDnsNamespace", data, meta)
}

func resourceServiceDiscoveryPublicDnsNamespaceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}