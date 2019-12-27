// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2EC2FleetType string = "AWS::EC2::EC2Fleet"

func DatasourceEC2EC2Fleet() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2EC2FleetRead,
		
		Schema: map[string]*schema.Schema{
			"target_capacity_specification": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetTargetCapacitySpecificationRequest(),
				Required: true,
				MaxItems: 1,
			},
			"on_demand_options": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetOnDemandOptionsRequest(),
				Optional: true,
				MaxItems: 1,
			},
			"type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"excess_capacity_termination_policy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tag_specifications": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetTagSpecification(),
				Optional: true,
			},
			"spot_options": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetSpotOptionsRequest(),
				Optional: true,
				MaxItems: 1,
			},
			"valid_from": {
				Type: schema.TypeString,
				Optional: true,
			},
			"replace_unhealthy_instances": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"launch_template_configs": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetFleetLaunchTemplateConfigRequest(),
				Required: true,
			},
			"terminate_instances_with_expiration": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"valid_until": {
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

func datasourceEC2EC2FleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2EC2FleetType, DatasourceEC2EC2Fleet(), data, meta)
}
