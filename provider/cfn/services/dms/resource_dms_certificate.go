// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html

package dms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const dMSCertificateType string = "AWS::DMS::Certificate"

func ResourceDMSCertificate() *schema.Resource {
	return &schema.Resource{
		Exists: resourceDMSCertificateExists,
		Read: resourceDMSCertificateRead,
		Create: resourceDMSCertificateCreate,
		Update: resourceDMSCertificateUpdate,
		Delete: resourceDMSCertificateDelete,
		CustomizeDiff: resourceDMSCertificateCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"certificate_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"certificate_pem": {
				Type: schema.TypeString,
				Optional: true,
			},
			"certificate_wallet": {
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

func resourceDMSCertificateExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceDMSCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dMSCertificateType, ResourceDMSCertificate(), data, meta)
}

func resourceDMSCertificateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(dMSCertificateType, ResourceDMSCertificate(), data, meta)
}

func resourceDMSCertificateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(dMSCertificateType, ResourceDMSCertificate(), data, meta)
}

func resourceDMSCertificateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(dMSCertificateType, data, meta)
}

func resourceDMSCertificateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(dMSCertificateType, data, meta)
}
