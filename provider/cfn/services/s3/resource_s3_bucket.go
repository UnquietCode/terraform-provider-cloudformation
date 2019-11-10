// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

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
			"accelerate_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketAccelerateConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"access_control": {
				Type: schema.TypeString,
				Required: false,
			},
			"analytics_configurations": {
				Type: schema.TypeSet,
				Elem: propertyBucketAnalyticsConfiguration(),
				Required: false,
			},
			"bucket_encryption": {
				Type: schema.TypeList,
				Elem: propertyBucketBucketEncryption(),
				Required: false,
				MaxItems: 1,
			},
			"bucket_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"cors_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketCorsConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"inventory_configurations": {
				Type: schema.TypeSet,
				Elem: propertyBucketInventoryConfiguration(),
				Required: false,
			},
			"lifecycle_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketLifecycleConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"logging_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketLoggingConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"metrics_configurations": {
				Type: schema.TypeSet,
				Elem: propertyBucketMetricsConfiguration(),
				Required: false,
			},
			"notification_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketNotificationConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"object_lock_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketObjectLockConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"object_lock_enabled": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"public_access_block_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketPublicAccessBlockConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"replication_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketReplicationConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"versioning_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketVersioningConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"website_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketWebsiteConfiguration(),
				Required: false,
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