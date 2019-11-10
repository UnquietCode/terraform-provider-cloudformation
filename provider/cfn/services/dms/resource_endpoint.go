// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package dms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceEndpointCreate,
		Read:   resourceEndpointRead,
		Update: resourceEndpointUpdate,
		Delete: resourceEndpointDelete,

		Schema: map[string]*schema.Schema{
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"port": {
				Type: schema.TypeInt,
				Required: false,
			},
			"database_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"elasticsearch_settings": {
				Type: schema.TypeList,
				Elem: propertyEndpointElasticsearchSettings(),
				Required: false,
				MaxItems: 1,
			},
			"s3_settings": {
				Type: schema.TypeList,
				Elem: propertyEndpointS3Settings(),
				Required: false,
				MaxItems: 1,
			},
			"engine_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"dynamo_db_settings": {
				Type: schema.TypeList,
				Elem: propertyEndpointDynamoDbSettings(),
				Required: false,
				MaxItems: 1,
			},
			"kinesis_settings": {
				Type: schema.TypeList,
				Elem: propertyEndpointKinesisSettings(),
				Required: false,
				MaxItems: 1,
			},
			"username": {
				Type: schema.TypeString,
				Required: false,
			},
			"ssl_mode": {
				Type: schema.TypeString,
				Required: false,
			},
			"server_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"extra_connection_attributes": {
				Type: schema.TypeString,
				Required: false,
			},
			"endpoint_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
				ForceNew: true,
			},
			"endpoint_identifier": {
				Type: schema.TypeString,
				Required: false,
			},
			"password": {
				Type: schema.TypeString,
				Required: false,
			},
			"certificate_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"mongo_db_settings": {
				Type: schema.TypeList,
				Elem: propertyEndpointMongoDbSettings(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DMS::Endpoint", data, meta)
}

func resourceEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DMS::Endpoint", data, meta)
}

func resourceEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DMS::Endpoint", data, meta)
}

func resourceEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DMS::Endpoint", data, meta)
}