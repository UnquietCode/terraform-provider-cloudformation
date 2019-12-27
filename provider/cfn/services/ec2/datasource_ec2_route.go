// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2RouteType string = "AWS::EC2::Route"

func DatasourceEC2Route() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2RouteRead,
		
		Schema: map[string]*schema.Schema{
			"destination_cidr_block": {
				Type: schema.TypeString,
				Optional: true,
			},
			"destination_ipv6_cidr_block": {
				Type: schema.TypeString,
				Optional: true,
			},
			"egress_only_internet_gateway_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"gateway_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"nat_gateway_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"route_table_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"transit_gateway_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"vpc_peering_connection_id": {
				Type: schema.TypeString,
				Optional: true,
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

func datasourceEC2RouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2RouteType, DatasourceEC2Route(), data, meta)
}
