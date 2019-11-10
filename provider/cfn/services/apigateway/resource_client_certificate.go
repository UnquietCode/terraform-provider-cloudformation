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

func ResourceClientCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceClientCertificateCreate,
		Read:   resourceClientCertificateRead,
		Update: resourceClientCertificateUpdate,
		Delete: resourceClientCertificateDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceClientCertificateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::ClientCertificate", data, meta)
}

func resourceClientCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::ClientCertificate", data, meta)
}

func resourceClientCertificateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::ClientCertificate", data, meta)
}

func resourceClientCertificateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::ClientCertificate", data, meta)
}