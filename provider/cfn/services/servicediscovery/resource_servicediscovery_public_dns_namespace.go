// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-publicdnsnamespace.html

package servicediscovery

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceDiscoveryPublicDnsNamespaceType string = "AWS::ServiceDiscovery::PublicDnsNamespace"

func ResourceServiceDiscoveryPublicDnsNamespace() *schema.Resource {
	return &schema.Resource{
		Read: resourceServiceDiscoveryPublicDnsNamespaceRead,
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

func resourceServiceDiscoveryPublicDnsNamespaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceDiscoveryPublicDnsNamespaceType, ResourceServiceDiscoveryPublicDnsNamespace(), data, meta)
}

func resourceServiceDiscoveryPublicDnsNamespaceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceDiscoveryPublicDnsNamespaceType, ResourceServiceDiscoveryPublicDnsNamespace(), data, meta)
}

func resourceServiceDiscoveryPublicDnsNamespaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceDiscoveryPublicDnsNamespaceType, ResourceServiceDiscoveryPublicDnsNamespace(), data, meta)
}

func resourceServiceDiscoveryPublicDnsNamespaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceDiscoveryPublicDnsNamespaceType, data, meta)
}

func resourceServiceDiscoveryPublicDnsNamespaceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceDiscoveryPublicDnsNamespaceType, data, meta)
}
