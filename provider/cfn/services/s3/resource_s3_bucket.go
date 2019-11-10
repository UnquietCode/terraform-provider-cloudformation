// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceS3Bucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceS3BucketCreate,
		Read:   resourceS3BucketRead,
		Update: resourceS3BucketUpdate,
		Delete: resourceS3BucketDelete,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"domain_name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"dual_stack_domain_name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"regional_domain_name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"website_url": {
				Type: schema.TypeString,
				Computed: true,
			},
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
				ForceNew: true,
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
				ForceNew: true,
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
		},
	}
}

func resourceS3BucketCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::S3::Bucket", data, meta)
}

func resourceS3BucketRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::S3::Bucket", data, meta)
}

func resourceS3BucketUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::S3::Bucket", data, meta)
}

func resourceS3BucketDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::S3::Bucket", data, meta)
}