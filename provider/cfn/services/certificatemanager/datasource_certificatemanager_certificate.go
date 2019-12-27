// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html

package certificatemanager

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const certificateManagerCertificateType string = "AWS::CertificateManager::Certificate"

func DatasourceCertificateManagerCertificate() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCertificateManagerCertificateRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceCertificateManagerCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(certificateManagerCertificateType, DatasourceCertificateManagerCertificate(), data, meta)
}
