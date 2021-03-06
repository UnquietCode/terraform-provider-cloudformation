// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html

package apigateway

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayRestApiType string = "AWS::ApiGateway::RestApi"

func DatasourceApiGatewayRestApi() *schema.Resource {
	return &schema.Resource{
		Read: datasourceApiGatewayRestApiRead,
		
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
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"body_s3_location": {
				Type: schema.TypeList,
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
				Type: schema.TypeList,
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
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"tags": misc.PropertyTags(),
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

func datasourceApiGatewayRestApiRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayRestApiType, DatasourceApiGatewayRestApi(), data, meta)
}
