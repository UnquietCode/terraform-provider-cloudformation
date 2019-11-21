// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html

package apigateway

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayDomainNameType string = "AWS::ApiGateway::DomainName"

func ResourceApiGatewayDomainName() *schema.Resource {
	return &schema.Resource{
		Read: resourceApiGatewayDomainNameRead,
		Create: resourceApiGatewayDomainNameCreate,
		Update: resourceApiGatewayDomainNameUpdate,
		Delete: resourceApiGatewayDomainNameDelete,
		CustomizeDiff: resourceApiGatewayDomainNameCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"certificate_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"endpoint_configuration": {
				Type: schema.TypeSet,
				Elem: propertyDomainNameEndpointConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"regional_certificate_arn": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceApiGatewayDomainNameRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayDomainNameType, ResourceApiGatewayDomainName(), data, meta)
}

func resourceApiGatewayDomainNameCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayDomainNameType, ResourceApiGatewayDomainName(), data, meta)
}

func resourceApiGatewayDomainNameUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayDomainNameType, ResourceApiGatewayDomainName(), data, meta)
}

func resourceApiGatewayDomainNameDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayDomainNameType, data, meta)
}

func resourceApiGatewayDomainNameCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayDomainNameType, data, meta)
}
