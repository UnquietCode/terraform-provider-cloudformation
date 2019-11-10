// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package appstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceFleet() *schema.Resource {
	return &schema.Resource{
		Create: resourceFleetCreate,
		Read:   resourceFleetRead,
		Update: resourceFleetUpdate,
		Delete: resourceFleetDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"compute_capacity": {
				Type: schema.TypeList,
				Elem: propertyFleetComputeCapacity(),
				Required: true,
				MaxItems: 1,
			},
			"vpc_config": {
				Type: schema.TypeList,
				Elem: propertyFleetVpcConfig(),
				Required: false,
				MaxItems: 1,
			},
			"fleet_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"enable_default_internet_access": {
				Type: schema.TypeBool,
				Required: false,
			},
			"domain_join_info": {
				Type: schema.TypeList,
				Elem: propertyFleetDomainJoinInfo(),
				Required: false,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"image_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"max_user_duration_in_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"idle_disconnect_timeout_in_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"disconnect_timeout_in_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"display_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"image_arn": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceFleetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppStream::Fleet", data, meta)
}

func resourceFleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppStream::Fleet", data, meta)
}

func resourceFleetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppStream::Fleet", data, meta)
}

func resourceFleetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppStream::Fleet", data, meta)
}