// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayClientCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayClientCertificateCreate,
		Read:   resourceApiGatewayClientCertificateRead,
		Update: resourceApiGatewayClientCertificateUpdate,
		Delete: resourceApiGatewayClientCertificateDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceApiGatewayClientCertificateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::ClientCertificate", data, meta)
}

func resourceApiGatewayClientCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::ClientCertificate", data, meta)
}

func resourceApiGatewayClientCertificateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::ClientCertificate", data, meta)
}

func resourceApiGatewayClientCertificateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::ClientCertificate", data, meta)
}