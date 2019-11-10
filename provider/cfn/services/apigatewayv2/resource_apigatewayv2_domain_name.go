// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayV2DomainName() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayV2DomainNameCreate,
		Read:   resourceApiGatewayV2DomainNameRead,
		Update: resourceApiGatewayV2DomainNameUpdate,
		Delete: resourceApiGatewayV2DomainNameDelete,

		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domain_name_configurations": {
				Type: schema.TypeList,
				Elem: propertyDomainNameDomainNameConfiguration(),
				Required: false,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceApiGatewayV2DomainNameCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::DomainName", data, meta)
}

func resourceApiGatewayV2DomainNameRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::DomainName", data, meta)
}

func resourceApiGatewayV2DomainNameUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::DomainName", data, meta)
}

func resourceApiGatewayV2DomainNameDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::DomainName", data, meta)
}