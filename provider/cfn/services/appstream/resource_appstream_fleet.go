// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html

package appstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppStreamFleet() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAppStreamFleetExists,
		Read: resourceAppStreamFleetRead,
		Create: resourceAppStreamFleetCreate,
		Update: resourceAppStreamFleetUpdate,
		Delete: resourceAppStreamFleetDelete,
		CustomizeDiff: resourceAppStreamFleetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
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
				Optional: true,
				MaxItems: 1,
			},
			"fleet_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"enable_default_internet_access": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"domain_join_info": {
				Type: schema.TypeList,
				Elem: propertyFleetDomainJoinInfo(),
				Optional: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"image_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"max_user_duration_in_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"idle_disconnect_timeout_in_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"disconnect_timeout_in_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"display_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"image_arn": {
				Type: schema.TypeString,
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

func resourceAppStreamFleetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAppStreamFleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppStream::Fleet", ResourceAppStreamFleet(), data, meta)
}

func resourceAppStreamFleetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppStream::Fleet", ResourceAppStreamFleet(), data, meta)
}

func resourceAppStreamFleetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppStream::Fleet", ResourceAppStreamFleet(), data, meta)
}

func resourceAppStreamFleetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppStream::Fleet", data, meta)
}

func resourceAppStreamFleetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
