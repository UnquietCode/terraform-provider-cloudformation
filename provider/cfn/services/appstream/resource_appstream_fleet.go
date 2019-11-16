// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

const appStreamFleetType string = "AWS::AppStream::Fleet"

var appStreamFleetProperties map[string]string = map[string]string{
	"description": "Description",
	"compute_capacity": "ComputeCapacity",
	"vpc_config": "VpcConfig",
	"fleet_type": "FleetType",
	"enable_default_internet_access": "EnableDefaultInternetAccess",
	"domain_join_info": "DomainJoinInfo",
	"name": "Name",
	"image_name": "ImageName",
	"max_user_duration_in_seconds": "MaxUserDurationInSeconds",
	"idle_disconnect_timeout_in_seconds": "IdleDisconnectTimeoutInSeconds",
	"disconnect_timeout_in_seconds": "DisconnectTimeoutInSeconds",
	"display_name": "DisplayName",
	"instance_type": "InstanceType",
	"tags": "Tags",
	"image_arn": "ImageArn",
}

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
			"tags": misc.PropertyTags(),
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
	return plugin.ResourceRead(appStreamFleetType, ResourceAppStreamFleet(), data, meta)
}

func resourceAppStreamFleetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appStreamFleetType, ResourceAppStreamFleet(), data, appStreamFleetProperties, meta)
}

func resourceAppStreamFleetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appStreamFleetType, ResourceAppStreamFleet(), data, appStreamFleetProperties, meta)
}

func resourceAppStreamFleetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appStreamFleetType, data, meta)
}

func resourceAppStreamFleetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appStreamFleetType, data, meta)
}
