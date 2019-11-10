// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

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
			"certificate_signing_request": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIoTCertificateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT::Certificate", data, meta)
}

func resourceIoTCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT::Certificate", data, meta)
}

func resourceIoTCertificateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT::Certificate", data, meta)
}

func resourceIoTCertificateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT::Certificate", data, meta)
}