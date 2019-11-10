// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2CapacityReservation() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2CapacityReservationCreate,
		Read:   resourceEC2CapacityReservationRead,
		Update: resourceEC2CapacityReservationUpdate,
		Delete: resourceEC2CapacityReservationDelete,

		Schema: map[string]*schema.Schema{
			"tenancy": {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"available_instance_count": {
				Type: schema.TypeInt,
				Computed: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: true,
				Computed: true,
				ForceNew: true,
			},
			"total_instance_count": {
				Type: schema.TypeInt,
				Computed: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
				Computed: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"instance_platform": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ephemeral_storage": {
				Type: schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"instance_match_criteria": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"end_date": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2CapacityReservationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::CapacityReservation", data, meta)
}

func resourceEC2CapacityReservationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::CapacityReservation", data, meta)
}

func resourceEC2CapacityReservationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::CapacityReservation", data, meta)
}

func resourceEC2CapacityReservationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::CapacityReservation", data, meta)
}