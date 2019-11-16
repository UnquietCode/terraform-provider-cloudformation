// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html

package s3

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const s3BucketType string = "AWS::S3::Bucket"

var s3BucketProperties map[string]string = map[string]string{
	"accelerate_configuration": "AccelerateConfiguration",
	"access_control": "AccessControl",
	"analytics_configurations": "AnalyticsConfigurations",
	"bucket_encryption": "BucketEncryption",
	"bucket_name": "BucketName",
	"cors_configuration": "CorsConfiguration",
	"inventory_configurations": "InventoryConfigurations",
	"lifecycle_configuration": "LifecycleConfiguration",
	"logging_configuration": "LoggingConfiguration",
	"metrics_configurations": "MetricsConfigurations",
	"notification_configuration": "NotificationConfiguration",
	"object_lock_configuration": "ObjectLockConfiguration",
	"object_lock_enabled": "ObjectLockEnabled",
	"public_access_block_configuration": "PublicAccessBlockConfiguration",
	"replication_configuration": "ReplicationConfiguration",
	"tags": "Tags",
	"versioning_configuration": "VersioningConfiguration",
	"website_configuration": "WebsiteConfiguration",
}

func ResourceS3Bucket() *schema.Resource {
	return &schema.Resource{
		Exists: resourceS3BucketExists,
		Read: resourceS3BucketRead,
		Create: resourceS3BucketCreate,
		Update: resourceS3BucketUpdate,
		Delete: resourceS3BucketDelete,
		CustomizeDiff: resourceS3BucketCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"accelerate_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketAccelerateConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"access_control": {
				Type: schema.TypeString,
				Optional: true,
			},
			"analytics_configurations": {
				Type: schema.TypeSet,
				Elem: propertyBucketAnalyticsConfiguration(),
				Optional: true,
			},
			"bucket_encryption": {
				Type: schema.TypeList,
				Elem: propertyBucketBucketEncryption(),
				Optional: true,
				MaxItems: 1,
			},
			"bucket_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cors_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketCorsConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"inventory_configurations": {
				Type: schema.TypeSet,
				Elem: propertyBucketInventoryConfiguration(),
				Optional: true,
			},
			"lifecycle_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketLifecycleConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"logging_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketLoggingConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"metrics_configurations": {
				Type: schema.TypeSet,
				Elem: propertyBucketMetricsConfiguration(),
				Optional: true,
			},
			"notification_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketNotificationConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"object_lock_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketObjectLockConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"object_lock_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"public_access_block_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketPublicAccessBlockConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"replication_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketReplicationConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"versioning_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketVersioningConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"website_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketWebsiteConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceS3BucketExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceS3BucketRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(s3BucketType, ResourceS3Bucket(), data, meta)
}

func resourceS3BucketCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(s3BucketType, ResourceS3Bucket(), data, s3BucketProperties, meta)
}

func resourceS3BucketUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(s3BucketType, ResourceS3Bucket(), data, s3BucketProperties, meta)
}

func resourceS3BucketDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(s3BucketType, data, meta)
}

func resourceS3BucketCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(s3BucketType, data, meta)
}
