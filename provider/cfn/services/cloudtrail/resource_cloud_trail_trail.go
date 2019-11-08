// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package cloudtrail

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCloudTrailTrail() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudTrailTrailCreate,
		Read:   resourceCloudTrailTrailRead,
		Update: resourceCloudTrailTrailUpdate,
		Delete: resourceCloudTrailTrailDelete,

		Schema: map[string]*schema.Schema{
			"cloud_watch_logs_log_group_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"cloud_watch_logs_role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"enable_log_file_validation": {
				Type: schema.TypeBool,
				Required: false,
			},
			"event_selectors": {
				Type: schema.TypeSet,
				Elem: propertyEventSelector(),
				Required: false,
			},
			"include_global_service_events": {
				Type: schema.TypeBool,
				Required: false,
			},
			"is_logging": {
				Type: schema.TypeBool,
				Required: true,
			},
			"is_multi_region_trail": {
				Type: schema.TypeBool,
				Required: false,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"s3_bucket_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"s3_key_prefix": {
				Type: schema.TypeString,
				Required: false,
			},
			"sns_topic_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"trail_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceCloudTrailTrailCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudTrail::Trail", data, meta)
}

func resourceCloudTrailTrailRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudTrail::Trail", data, meta)
}

func resourceCloudTrailTrailUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudTrail::Trail", data, meta)
}

func resourceCloudTrailTrailDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudTrail::Trail", data, meta)
}