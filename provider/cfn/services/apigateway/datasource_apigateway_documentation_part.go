// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html

package apigateway

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayDocumentationPartType string = "AWS::ApiGateway::DocumentationPart"

func DatasourceApiGatewayDocumentationPart() *schema.Resource {
	return &schema.Resource{
		Read: datasourceApiGatewayDocumentationPartRead,
		
		Schema: map[string]*schema.Schema{
			"location": {
				Type: schema.TypeList,
				Elem: propertyDocumentationPartLocation(),
				Required: true,
				MaxItems: 1,
			},
			"properties": {
				Type: schema.TypeString,
				Required: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
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

func datasourceApiGatewayDocumentationPartRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayDocumentationPartType, DatasourceApiGatewayDocumentationPart(), data, meta)
}
