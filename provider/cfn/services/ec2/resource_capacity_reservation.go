// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCapacityReservation() *schema.Resource {
	return &schema.Resource{
		Create: resourceCapacityReservationCreate,
		Read:   resourceCapacityReservationRead,
		Update: resourceCapacityReservationUpdate,
		Delete: resourceCapacityReservationDelete,

		Schema: map[string]*schema.Schema{
			"tenancy": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"end_date_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"instance_count": {
				Type: schema.TypeInt,
				Required: true,
			},
			"tag_specifications": {
				Type: schema.TypeList,
				Elem: propertyCapacityReservationTagSpecification(),
				Required: false,
				ForceNew: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_platform": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ephemeral_storage": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"instance_match_criteria": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"end_date": {
				Type: schema.TypeString,
				Required: false,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceCapacityReservationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::CapacityReservation", data, meta)
}

func resourceCapacityReservationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::CapacityReservation", data, meta)
}

func resourceCapacityReservationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::CapacityReservation", data, meta)
}

func resourceCapacityReservationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::CapacityReservation", data, meta)
}