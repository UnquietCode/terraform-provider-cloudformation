// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayDomainName() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayDomainNameCreate,
		Read:   resourceApiGatewayDomainNameRead,
		Update: resourceApiGatewayDomainNameUpdate,
		Delete: resourceApiGatewayDomainNameDelete,

		Schema: map[string]*schema.Schema{
			"distribution_domain_name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"distribution_hosted_zone_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"regional_domain_name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"regional_hosted_zone_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"certificate_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"endpoint_configuration": {
				Type: schema.TypeList,
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
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApiGatewayDomainNameCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::DomainName", ResourceApiGatewayDomainName(), data, meta)
}

func resourceApiGatewayDomainNameRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::DomainName", ResourceApiGatewayDomainName(), data, meta)
}

func resourceApiGatewayDomainNameUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::DomainName", ResourceApiGatewayDomainName(), data, meta)
}

func resourceApiGatewayDomainNameDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::DomainName", data, meta)
}