// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html

package s3

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const s3BucketType string = "AWS::S3::Bucket"

func DatasourceS3Bucket() *schema.Resource {
	return &schema.Resource{
		Read: datasourceS3BucketRead,
		
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
			"tags": misc.PropertyTags(),
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceS3BucketRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(s3BucketType, DatasourceS3Bucket(), data, meta)
}
