// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html

package dms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDMSEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceDMSEndpointCreate,
		Read:   resourceDMSEndpointRead,
		Update: resourceDMSEndpointUpdate,
		Delete: resourceDMSEndpointDelete,

		Schema: map[string]*schema.Schema{
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
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
				Type: schema.TypeList,
				Elem: propertyEndpointElasticsearchSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"s3_settings": {
				Type: schema.TypeList,
				Elem: propertyEndpointS3Settings(),
				Optional: true,
				MaxItems: 1,
			},
			"engine_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"dynamo_db_settings": {
				Type: schema.TypeList,
				Elem: propertyEndpointDynamoDbSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"kinesis_settings": {
				Type: schema.TypeList,
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
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
				ForceNew: true,
			},
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
				Type: schema.TypeList,
				Elem: propertyEndpointMongoDbSettings(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceDMSEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DMS::Endpoint", data, meta)
}

func resourceDMSEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DMS::Endpoint", data, meta)
}

func resourceDMSEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DMS::Endpoint", data, meta)
}

func resourceDMSEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DMS::Endpoint", data, meta)
}