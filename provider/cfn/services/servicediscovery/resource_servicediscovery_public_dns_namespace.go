// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

var serviceDiscoveryPublicDnsNamespaceProperties map[string]string = map[string]string{
	"description": "Description",
	"name": "Name",
}

func ResourceServiceDiscoveryPublicDnsNamespace() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceDiscoveryPublicDnsNamespaceExists,
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

func resourceServiceDiscoveryPublicDnsNamespaceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceDiscoveryPublicDnsNamespaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceDiscoveryPublicDnsNamespaceType, ResourceServiceDiscoveryPublicDnsNamespace(), data, meta)
}

func resourceServiceDiscoveryPublicDnsNamespaceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceDiscoveryPublicDnsNamespaceType, ResourceServiceDiscoveryPublicDnsNamespace(), data, serviceDiscoveryPublicDnsNamespaceProperties, meta)
}

func resourceServiceDiscoveryPublicDnsNamespaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceDiscoveryPublicDnsNamespaceType, ResourceServiceDiscoveryPublicDnsNamespace(), data, serviceDiscoveryPublicDnsNamespaceProperties, meta)
}

func resourceServiceDiscoveryPublicDnsNamespaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceDiscoveryPublicDnsNamespaceType, data, meta)
}

func resourceServiceDiscoveryPublicDnsNamespaceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceDiscoveryPublicDnsNamespaceType, data, meta)
}
