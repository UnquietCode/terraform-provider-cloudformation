// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificate.html

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTCertificateType string = "AWS::IoT::Certificate"

func ResourceIoTCertificate() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIoTCertificateExists,
		Read: resourceIoTCertificateRead,
		Create: resourceIoTCertificateCreate,
		Update: resourceIoTCertificateUpdate,
		Delete: resourceIoTCertificateDelete,
		CustomizeDiff: resourceIoTCertificateCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"certificate_signing_request": {
				Type: schema.TypeString,
				Required: true,
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

func resourceIoTCertificateExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIoTCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTCertificateType, ResourceIoTCertificate(), data, meta)
}

func resourceIoTCertificateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTCertificateType, ResourceIoTCertificate(), data, meta)
}

func resourceIoTCertificateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTCertificateType, ResourceIoTCertificate(), data, meta)
}

func resourceIoTCertificateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTCertificateType, data, meta)
}

func resourceIoTCertificateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTCertificateType, data, meta)
}
