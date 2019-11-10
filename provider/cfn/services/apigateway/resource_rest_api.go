// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRestApi() *schema.Resource {
	return &schema.Resource{
		Create: resourceRestApiCreate,
		Read:   resourceRestApiRead,
		Update: resourceRestApiUpdate,
		Delete: resourceRestApiDelete,

		Schema: map[string]*schema.Schema{
			"api_key_source_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"binary_media_types": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"body": {
				Type: schema.TypeMap,
				Required: false,
			},
			"body_s3_location": {
				Type: schema.TypeList,
				Elem: propertyRestApiS3Location(),
				Required: false,
				MaxItems: 1,
			},
			"clone_from": {
				Type: schema.TypeString,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"endpoint_configuration": {
				Type: schema.TypeList,
				Elem: propertyRestApiEndpointConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"fail_on_warnings": {
				Type: schema.TypeBool,
				Required: false,
			},
			"minimum_compression_size": {
				Type: schema.TypeInt,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
			"parameters": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"policy": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceRestApiCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::RestApi", data, meta)
}

func resourceRestApiRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::RestApi", data, meta)
}

func resourceRestApiUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::RestApi", data, meta)
}

func resourceRestApiDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::RestApi", data, meta)
}