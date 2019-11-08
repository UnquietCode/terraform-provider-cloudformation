// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package dms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDMSCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceDMSCertificateCreate,
		Read:   resourceDMSCertificateRead,
		Update: resourceDMSCertificateUpdate,
		Delete: resourceDMSCertificateDelete,

		Schema: map[string]*schema.Schema{
			"certificate_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"certificate_pem": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"certificate_wallet": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceDMSCertificateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DMS::Certificate", data, meta)
}

func resourceDMSCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DMS::Certificate", data, meta)
}

func resourceDMSCertificateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DMS::Certificate", data, meta)
}

func resourceDMSCertificateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DMS::Certificate", data, meta)
}