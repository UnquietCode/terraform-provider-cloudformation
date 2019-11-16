// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayRestApiType string = "AWS::ApiGateway::RestApi"

func ResourceApiGatewayRestApi() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayRestApiExists,
		Read: resourceApiGatewayRestApiRead,
		Create: resourceApiGatewayRestApiCreate,
		Update: resourceApiGatewayRestApiUpdate,
		Delete: resourceApiGatewayRestApiDelete,
		CustomizeDiff: resourceApiGatewayRestApiCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"api_key_source_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"binary_media_types": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"body": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"body_s3_location": {
				Type: schema.TypeSet,
				Elem: propertyRestApiS3Location(),
				Optional: true,
				MaxItems: 1,
			},
			"clone_from": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"endpoint_configuration": {
				Type: schema.TypeSet,
				Elem: propertyRestApiEndpointConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"fail_on_warnings": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"minimum_compression_size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"parameters": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"policy": {
				Type: schema.TypeMap,
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

func resourceApiGatewayRestApiExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayRestApiRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayRestApiType, ResourceApiGatewayRestApi(), data, meta)
}

func resourceApiGatewayRestApiCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayRestApiType, ResourceApiGatewayRestApi(), data, meta)
}

func resourceApiGatewayRestApiUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayRestApiType, ResourceApiGatewayRestApi(), data, meta)
}

func resourceApiGatewayRestApiDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayRestApiType, data, meta)
}

func resourceApiGatewayRestApiCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayRestApiType, data, meta)
}
