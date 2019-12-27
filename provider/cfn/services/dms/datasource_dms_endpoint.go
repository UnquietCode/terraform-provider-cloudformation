// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceDMSEndpoint() *schema.Resource {
	return &schema.Resource{
		Read: datasourceDMSEndpointRead,
		
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
				Type: schema.TypeList,
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
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceDMSEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dMSEndpointType, DatasourceDMSEndpoint(), data, meta)
}
