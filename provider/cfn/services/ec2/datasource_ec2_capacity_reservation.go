// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2CapacityReservationType string = "AWS::EC2::CapacityReservation"

func DatasourceEC2CapacityReservation() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2CapacityReservationRead,
		
		Schema: map[string]*schema.Schema{
			"tenancy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"end_date_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_count": {
				Type: schema.TypeInt,
				Required: true,
			},
			"tag_specifications": {
				Type: schema.TypeList,
				Elem: propertyCapacityReservationTagSpecification(),
				Optional: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: true,
			},
			"instance_platform": {
				Type: schema.TypeString,
				Required: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"ephemeral_storage": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"instance_match_criteria": {
				Type: schema.TypeString,
				Optional: true,
			},
			"end_date": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
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

func datasourceEC2CapacityReservationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2CapacityReservationType, DatasourceEC2CapacityReservation(), data, meta)
}
