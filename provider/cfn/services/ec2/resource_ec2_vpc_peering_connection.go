// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VPCPeeringConnectionType string = "AWS::EC2::VPCPeeringConnection"

func ResourceEC2VPCPeeringConnection() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2VPCPeeringConnectionRead,
		Create: resourceEC2VPCPeeringConnectionCreate,
		Update: resourceEC2VPCPeeringConnectionUpdate,
		Delete: resourceEC2VPCPeeringConnectionDelete,
		CustomizeDiff: resourceEC2VPCPeeringConnectionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"peer_owner_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"peer_region": {
				Type: schema.TypeString,
				Optional: true,
			},
			"peer_role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"peer_vpc_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"vpc_id": {
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

func resourceEC2VPCPeeringConnectionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPCPeeringConnectionType, ResourceEC2VPCPeeringConnection(), data, meta)
}

func resourceEC2VPCPeeringConnectionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2VPCPeeringConnectionType, ResourceEC2VPCPeeringConnection(), data, meta)
}

func resourceEC2VPCPeeringConnectionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2VPCPeeringConnectionType, ResourceEC2VPCPeeringConnection(), data, meta)
}

func resourceEC2VPCPeeringConnectionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2VPCPeeringConnectionType, data, meta)
}

func resourceEC2VPCPeeringConnectionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2VPCPeeringConnectionType, data, meta)
}
