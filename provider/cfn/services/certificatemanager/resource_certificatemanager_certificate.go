// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceCertificateManagerCertificate() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCertificateManagerCertificateExists,
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
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
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

func resourceCertificateManagerCertificateExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCertificateManagerCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CertificateManager::Certificate", ResourceCertificateManagerCertificate(), data, meta)
}

func resourceCertificateManagerCertificateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CertificateManager::Certificate", ResourceCertificateManagerCertificate(), data, meta)
}

func resourceCertificateManagerCertificateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CertificateManager::Certificate", ResourceCertificateManagerCertificate(), data, meta)
}

func resourceCertificateManagerCertificateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CertificateManager::Certificate", data, meta)
}

func resourceCertificateManagerCertificateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::CertificateManager::Certificate", data, meta)
}

