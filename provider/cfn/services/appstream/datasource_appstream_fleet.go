// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html

package appstream

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appStreamFleetType string = "AWS::AppStream::Fleet"

func DatasourceAppStreamFleet() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAppStreamFleetRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceAppStreamFleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appStreamFleetType, DatasourceAppStreamFleet(), data, meta)
}
