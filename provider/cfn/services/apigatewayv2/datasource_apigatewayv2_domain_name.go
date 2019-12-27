// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html

package apigatewayv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayV2DomainNameType string = "AWS::ApiGatewayV2::DomainName"

func DatasourceApiGatewayV2DomainName() *schema.Resource {
	return &schema.Resource{
		Read: datasourceApiGatewayV2DomainNameRead,
		
		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"domain_name_configurations": {
				Type: schema.TypeList,
				Elem: propertyDomainNameDomainNameConfiguration(),
				Optional: true,
			},
			"tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
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

func datasourceApiGatewayV2DomainNameRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayV2DomainNameType, DatasourceApiGatewayV2DomainName(), data, meta)
}
