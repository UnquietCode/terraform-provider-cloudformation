// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html

package certificatemanager

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const certificateManagerCertificateType string = "AWS::CertificateManager::Certificate"

func ResourceCertificateManagerCertificate() *schema.Resource {
	return &schema.Resource{
		Read: resourceCertificateManagerCertificateRead,
		Create: resourceCertificateManagerCertificateCreate,
		Update: resourceCertificateManagerCertificateUpdate,
		Delete: resourceCertificateManagerCertificateDelete,
		CustomizeDiff: resourceCertificateManagerCertificateCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"domain_validation_options": {
				Type: schema.TypeSet,
				Elem: propertyCertificateDomainValidationOption(),
				Optional: true,
			},
			"subject_alternative_names": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"tags": misc.PropertyTags(),
			"validation_method": {
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

func resourceCertificateManagerCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(certificateManagerCertificateType, ResourceCertificateManagerCertificate(), data, meta)
}

func resourceCertificateManagerCertificateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(certificateManagerCertificateType, ResourceCertificateManagerCertificate(), data, meta)
}

func resourceCertificateManagerCertificateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(certificateManagerCertificateType, ResourceCertificateManagerCertificate(), data, meta)
}

func resourceCertificateManagerCertificateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(certificateManagerCertificateType, data, meta)
}

func resourceCertificateManagerCertificateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(certificateManagerCertificateType, data, meta)
}
