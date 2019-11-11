// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificate.html

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoTCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoTCertificateCreate,
		Read:   resourceIoTCertificateRead,
		Update: resourceIoTCertificateUpdate,
		Delete: resourceIoTCertificateDelete,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"certificate_signing_request": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIoTCertificateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT::Certificate", ResourceIoTCertificate(), data, meta)
}

func resourceIoTCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT::Certificate", ResourceIoTCertificate(), data, meta)
}

func resourceIoTCertificateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT::Certificate", ResourceIoTCertificate(), data, meta)
}

func resourceIoTCertificateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT::Certificate", data, meta)
}