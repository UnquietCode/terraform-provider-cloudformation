// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservice.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VPCEndpointServiceType string = "AWS::EC2::VPCEndpointService"

func ResourceEC2VPCEndpointService() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2VPCEndpointServiceRead,
		Create: resourceEC2VPCEndpointServiceCreate,
		Update: resourceEC2VPCEndpointServiceUpdate,
		Delete: resourceEC2VPCEndpointServiceDelete,
		CustomizeDiff: resourceEC2VPCEndpointServiceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"network_load_balancer_arns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"acceptance_required": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2VPCEndpointServiceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPCEndpointServiceType, ResourceEC2VPCEndpointService(), data, meta)
}

func resourceEC2VPCEndpointServiceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2VPCEndpointServiceType, ResourceEC2VPCEndpointService(), data, meta)
}

func resourceEC2VPCEndpointServiceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2VPCEndpointServiceType, ResourceEC2VPCEndpointService(), data, meta)
}

func resourceEC2VPCEndpointServiceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2VPCEndpointServiceType, data, meta)
}

func resourceEC2VPCEndpointServiceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2VPCEndpointServiceType, data, meta)
}
