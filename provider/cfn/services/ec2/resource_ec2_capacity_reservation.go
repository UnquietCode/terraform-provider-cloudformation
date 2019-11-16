// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2CapacityReservationType string = "AWS::EC2::CapacityReservation"

func ResourceEC2CapacityReservation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2CapacityReservationExists,
		Read: resourceEC2CapacityReservationRead,
		Create: resourceEC2CapacityReservationCreate,
		Update: resourceEC2CapacityReservationUpdate,
		Delete: resourceEC2CapacityReservationDelete,
		CustomizeDiff: resourceEC2CapacityReservationCustomizeDiff,
		
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
			},
		},
	}
}

func resourceEC2CapacityReservationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2CapacityReservationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2CapacityReservationType, ResourceEC2CapacityReservation(), data, meta)
}

func resourceEC2CapacityReservationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2CapacityReservationType, ResourceEC2CapacityReservation(), data, meta)
}

func resourceEC2CapacityReservationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2CapacityReservationType, ResourceEC2CapacityReservation(), data, meta)
}

func resourceEC2CapacityReservationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2CapacityReservationType, data, meta)
}

func resourceEC2CapacityReservationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2CapacityReservationType, data, meta)
}
