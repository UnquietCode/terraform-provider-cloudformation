// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceCertificateManagerCertificateCreate,
		Read:   resourceCertificateManagerCertificateRead,
		Update: resourceCertificateManagerCertificateUpdate,
		Delete: resourceCertificateManagerCertificateDelete,

		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domain_validation_options": {
				Type: schema.TypeSet,
				Elem: propertyCertificateDomainValidationOption(),
				Required: false,
				ForceNew: true,
			},
			"subject_alternative_names": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"validation_method": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceCertificateManagerCertificateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CertificateManager::Certificate", data, meta)
}

func resourceCertificateManagerCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CertificateManager::Certificate", data, meta)
}

func resourceCertificateManagerCertificateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CertificateManager::Certificate", data, meta)
}

func resourceCertificateManagerCertificateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CertificateManager::Certificate", data, meta)
}