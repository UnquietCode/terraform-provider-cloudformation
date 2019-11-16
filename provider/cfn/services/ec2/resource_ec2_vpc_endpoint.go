// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VPCEndpointType string = "AWS::EC2::VPCEndpoint"

var eC2VPCEndpointProperties map[string]string = map[string]string{
	"policy_document": "PolicyDocument",
	"private_dns_enabled": "PrivateDnsEnabled",
	"route_table_ids": "RouteTableIds",
	"security_group_ids": "SecurityGroupIds",
	"service_name": "ServiceName",
	"subnet_ids": "SubnetIds",
	"vpc_endpoint_type": "VpcEndpointType",
	"vpc_id": "VpcId",
}

func ResourceEC2VPCEndpoint() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2VPCEndpointExists,
		Read: resourceEC2VPCEndpointRead,
		Create: resourceEC2VPCEndpointCreate,
		Update: resourceEC2VPCEndpointUpdate,
		Delete: resourceEC2VPCEndpointDelete,
		CustomizeDiff: resourceEC2VPCEndpointCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"policy_document": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"private_dns_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"route_table_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"security_group_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"service_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"subnet_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"vpc_endpoint_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"vpc_id": {
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

func resourceEC2VPCEndpointExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2VPCEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPCEndpointType, ResourceEC2VPCEndpoint(), data, meta)
}

func resourceEC2VPCEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2VPCEndpointType, ResourceEC2VPCEndpoint(), data, eC2VPCEndpointProperties, meta)
}

func resourceEC2VPCEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2VPCEndpointType, ResourceEC2VPCEndpoint(), data, eC2VPCEndpointProperties, meta)
}

func resourceEC2VPCEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2VPCEndpointType, data, meta)
}

func resourceEC2VPCEndpointCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2VPCEndpointType, data, meta)
}
