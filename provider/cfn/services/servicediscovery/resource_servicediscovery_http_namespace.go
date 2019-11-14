// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-httpnamespace.html

package servicediscovery

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceDiscoveryHttpNamespace() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceDiscoveryHttpNamespaceExists,
		Read: resourceServiceDiscoveryHttpNamespaceRead,
		Create: resourceServiceDiscoveryHttpNamespaceCreate,
		Update: resourceServiceDiscoveryHttpNamespaceUpdate,
		Delete: resourceServiceDiscoveryHttpNamespaceDelete,
		CustomizeDiff: resourceServiceDiscoveryHttpNamespaceCustomizeDiff,
		
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

func resourceServiceDiscoveryHttpNamespaceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceDiscoveryHttpNamespaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceDiscovery::HttpNamespace", ResourceServiceDiscoveryHttpNamespace(), data, meta)
}

func resourceServiceDiscoveryHttpNamespaceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceDiscovery::HttpNamespace", ResourceServiceDiscoveryHttpNamespace(), data, meta)
}

func resourceServiceDiscoveryHttpNamespaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceDiscovery::HttpNamespace", ResourceServiceDiscoveryHttpNamespace(), data, meta)
}

func resourceServiceDiscoveryHttpNamespaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceDiscovery::HttpNamespace", data, meta)
}

func resourceServiceDiscoveryHttpNamespaceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
