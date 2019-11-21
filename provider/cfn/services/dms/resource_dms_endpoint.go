// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html

package dms

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const dMSEndpointType string = "AWS::DMS::Endpoint"

func ResourceDMSEndpoint() *schema.Resource {
	return &schema.Resource{
		Read: resourceDMSEndpointRead,
		Create: resourceDMSEndpointCreate,
		Update: resourceDMSEndpointUpdate,
		Delete: resourceDMSEndpointDelete,
		CustomizeDiff: resourceDMSEndpointCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"port": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"database_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"elasticsearch_settings": {
				Type: schema.TypeSet,
				Elem: propertyEndpointElasticsearchSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"s3_settings": {
				Type: schema.TypeSet,
				Elem: propertyEndpointS3Settings(),
				Optional: true,
				MaxItems: 1,
			},
			"engine_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"dynamo_db_settings": {
				Type: schema.TypeSet,
				Elem: propertyEndpointDynamoDbSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"kinesis_settings": {
				Type: schema.TypeSet,
				Elem: propertyEndpointKinesisSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"username": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ssl_mode": {
				Type: schema.TypeString,
				Optional: true,
			},
			"server_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"extra_connection_attributes": {
				Type: schema.TypeString,
				Optional: true,
			},
			"endpoint_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"endpoint_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"password": {
				Type: schema.TypeString,
				Optional: true,
			},
			"certificate_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"mongo_db_settings": {
				Type: schema.TypeSet,
				Elem: propertyEndpointMongoDbSettings(),
				Optional: true,
				MaxItems: 1,
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

func resourceDMSEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dMSEndpointType, ResourceDMSEndpoint(), data, meta)
}

func resourceDMSEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(dMSEndpointType, ResourceDMSEndpoint(), data, meta)
}

func resourceDMSEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(dMSEndpointType, ResourceDMSEndpoint(), data, meta)
}

func resourceDMSEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(dMSEndpointType, data, meta)
}

func resourceDMSEndpointCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(dMSEndpointType, data, meta)
}
