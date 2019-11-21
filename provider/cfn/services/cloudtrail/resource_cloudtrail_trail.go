// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html

package cloudtrail

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloudTrailTrailType string = "AWS::CloudTrail::Trail"

func ResourceCloudTrailTrail() *schema.Resource {
	return &schema.Resource{
		Read: resourceCloudTrailTrailRead,
		Create: resourceCloudTrailTrailCreate,
		Update: resourceCloudTrailTrailUpdate,
		Delete: resourceCloudTrailTrailDelete,
		CustomizeDiff: resourceCloudTrailTrailCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"cloud_watch_logs_log_group_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cloud_watch_logs_role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"enable_log_file_validation": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"event_selectors": {
				Type: schema.TypeSet,
				Elem: propertyTrailEventSelector(),
				Optional: true,
			},
			"include_global_service_events": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"is_logging": {
				Type: schema.TypeBool,
				Required: true,
			},
			"is_multi_region_trail": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"s3_bucket_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"s3_key_prefix": {
				Type: schema.TypeString,
				Optional: true,
			},
			"sns_topic_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"trail_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceCloudTrailTrailRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloudTrailTrailType, ResourceCloudTrailTrail(), data, meta)
}

func resourceCloudTrailTrailCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cloudTrailTrailType, ResourceCloudTrailTrail(), data, meta)
}

func resourceCloudTrailTrailUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cloudTrailTrailType, ResourceCloudTrailTrail(), data, meta)
}

func resourceCloudTrailTrailDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cloudTrailTrailType, data, meta)
}

func resourceCloudTrailTrailCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cloudTrailTrailType, data, meta)
}
