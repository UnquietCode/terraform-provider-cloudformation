// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html

package cloudtrail

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloudTrailTrailType string = "AWS::CloudTrail::Trail"

var cloudTrailTrailProperties map[string]string = map[string]string{
	"cloud_watch_logs_log_group_arn": "CloudWatchLogsLogGroupArn",
	"cloud_watch_logs_role_arn": "CloudWatchLogsRoleArn",
	"enable_log_file_validation": "EnableLogFileValidation",
	"event_selectors": "EventSelectors",
	"include_global_service_events": "IncludeGlobalServiceEvents",
	"is_logging": "IsLogging",
	"is_multi_region_trail": "IsMultiRegionTrail",
	"kms_key_id": "KMSKeyId",
	"s3_bucket_name": "S3BucketName",
	"s3_key_prefix": "S3KeyPrefix",
	"sns_topic_name": "SnsTopicName",
	"tags": "Tags",
	"trail_name": "TrailName",
}

func ResourceCloudTrailTrail() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCloudTrailTrailExists,
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
			},
		},
	}
}

func resourceCloudTrailTrailExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCloudTrailTrailRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloudTrailTrailType, ResourceCloudTrailTrail(), data, meta)
}

func resourceCloudTrailTrailCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cloudTrailTrailType, ResourceCloudTrailTrail(), data, cloudTrailTrailProperties, meta)
}

func resourceCloudTrailTrailUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cloudTrailTrailType, ResourceCloudTrailTrail(), data, cloudTrailTrailProperties, meta)
}

func resourceCloudTrailTrailDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cloudTrailTrailType, data, meta)
}

func resourceCloudTrailTrailCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cloudTrailTrailType, data, meta)
}
