// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html

package amplify

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const amplifyDomainType string = "AWS::Amplify::Domain"

func ResourceAmplifyDomain() *schema.Resource {
	return &schema.Resource{
		Read: resourceAmplifyDomainRead,
		Create: resourceAmplifyDomainCreate,
		Update: resourceAmplifyDomainUpdate,
		Delete: resourceAmplifyDomainDelete,
		CustomizeDiff: resourceAmplifyDomainCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"sub_domain_settings": {
				Type: schema.TypeList,
				Elem: propertyDomainSubDomainSetting(),
				Required: true,
			},
			"app_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceAmplifyDomainRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(amplifyDomainType, ResourceAmplifyDomain(), data, meta)
}

func resourceAmplifyDomainCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(amplifyDomainType, ResourceAmplifyDomain(), data, meta)
}

func resourceAmplifyDomainUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(amplifyDomainType, ResourceAmplifyDomain(), data, meta)
}

func resourceAmplifyDomainDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(amplifyDomainType, data, meta)
}

func resourceAmplifyDomainCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(amplifyDomainType, data, meta)
}
