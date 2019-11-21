// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-instance.html

package servicediscovery

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceDiscoveryInstanceType string = "AWS::ServiceDiscovery::Instance"

func ResourceServiceDiscoveryInstance() *schema.Resource {
	return &schema.Resource{
		Read: resourceServiceDiscoveryInstanceRead,
		Create: resourceServiceDiscoveryInstanceCreate,
		Update: resourceServiceDiscoveryInstanceUpdate,
		Delete: resourceServiceDiscoveryInstanceDelete,
		CustomizeDiff: resourceServiceDiscoveryInstanceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"instance_attributes": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"instance_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"service_id": {
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

func resourceServiceDiscoveryInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceDiscoveryInstanceType, ResourceServiceDiscoveryInstance(), data, meta)
}

func resourceServiceDiscoveryInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceDiscoveryInstanceType, ResourceServiceDiscoveryInstance(), data, meta)
}

func resourceServiceDiscoveryInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceDiscoveryInstanceType, ResourceServiceDiscoveryInstance(), data, meta)
}

func resourceServiceDiscoveryInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceDiscoveryInstanceType, data, meta)
}

func resourceServiceDiscoveryInstanceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceDiscoveryInstanceType, data, meta)
}
