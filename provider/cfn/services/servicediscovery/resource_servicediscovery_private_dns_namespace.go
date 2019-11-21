// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-privatednsnamespace.html

package servicediscovery

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceDiscoveryPrivateDnsNamespaceType string = "AWS::ServiceDiscovery::PrivateDnsNamespace"

func ResourceServiceDiscoveryPrivateDnsNamespace() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceServiceDiscoveryPrivateDnsNamespaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceDiscoveryPrivateDnsNamespaceType, ResourceServiceDiscoveryPrivateDnsNamespace(), data, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceDiscoveryPrivateDnsNamespaceType, ResourceServiceDiscoveryPrivateDnsNamespace(), data, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceDiscoveryPrivateDnsNamespaceType, ResourceServiceDiscoveryPrivateDnsNamespace(), data, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceDiscoveryPrivateDnsNamespaceType, data, meta)
}

func resourceServiceDiscoveryPrivateDnsNamespaceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceDiscoveryPrivateDnsNamespaceType, data, meta)
}
