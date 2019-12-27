// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceEC2VPCPeeringConnection() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2VPCPeeringConnectionRead,
		
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
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceEC2VPCPeeringConnectionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPCPeeringConnectionType, DatasourceEC2VPCPeeringConnection(), data, meta)
}
