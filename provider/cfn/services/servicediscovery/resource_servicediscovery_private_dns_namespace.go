// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-privatednsnamespace.html

package servicediscovery

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceDiscoveryPrivateDnsNamespaceType string = "AWS::ServiceDiscovery::PrivateDnsNamespace"

var serviceDiscoveryPrivateDnsNamespaceProperties map[string]string = map[string]string{
	"description": "Description",
	"vpc": "Vpc",
	"name": "Name",
}

func ResourceServiceDiscoveryPrivateDnsNamespace() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceDiscoveryPrivateDnsNamespaceExists,
		Read: resourceServiceDiscoveryPrivateDnsNamespaceRead,
		Create: resourceServiceDiscoveryPrivateDnsNamespaceCreate,
		Update: resourceServiceDiscoveryPrivateDnsNamespaceUpdate,
		Delete: resourceServiceDiscoveryPrivateDnsNamespaceDelete,
		CustomizeDiff: resourceServiceDiscoveryPrivateDnsNamespaceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"vpc": {
				Type: schema.TypeString,
				Required: true,
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

func resourceServiceDiscoveryPrivateDnsNamespaceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceDiscoveryPrivateDnsNamespaceType, ResourceServiceDiscoveryPrivateDnsNamespace(), data, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceDiscoveryPrivateDnsNamespaceType, ResourceServiceDiscoveryPrivateDnsNamespace(), data, serviceDiscoveryPrivateDnsNamespaceProperties, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceDiscoveryPrivateDnsNamespaceType, ResourceServiceDiscoveryPrivateDnsNamespace(), data, serviceDiscoveryPrivateDnsNamespaceProperties, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceDiscoveryPrivateDnsNamespaceType, data, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceDiscoveryPrivateDnsNamespaceType, data, meta)
}
