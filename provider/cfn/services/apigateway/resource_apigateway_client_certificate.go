// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-clientcertificate.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayClientCertificateType string = "AWS::ApiGateway::ClientCertificate"

func ResourceApiGatewayClientCertificate() *schema.Resource {
	return &schema.Resource{
		Read: resourceApiGatewayClientCertificateRead,
		Create: resourceApiGatewayClientCertificateCreate,
		Update: resourceApiGatewayClientCertificateUpdate,
		Delete: resourceApiGatewayClientCertificateDelete,
		CustomizeDiff: resourceApiGatewayClientCertificateCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
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

func resourceApiGatewayClientCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayClientCertificateType, ResourceApiGatewayClientCertificate(), data, meta)
}

func resourceApiGatewayClientCertificateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayClientCertificateType, ResourceApiGatewayClientCertificate(), data, meta)
}

func resourceApiGatewayClientCertificateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayClientCertificateType, ResourceApiGatewayClientCertificate(), data, meta)
}

func resourceApiGatewayClientCertificateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayClientCertificateType, data, meta)
}

func resourceApiGatewayClientCertificateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayClientCertificateType, data, meta)
}
