// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package amplify

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAmplifyDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceAmplifyDomainCreate,
		Read:   resourceAmplifyDomainRead,
		Update: resourceAmplifyDomainUpdate,
		Delete: resourceAmplifyDomainDelete,

		Schema: map[string]*schema.Schema{
			"sub_domain_settings": {
				Type: schema.TypeList,
				Elem: propertyDomainSubDomainSetting(),
				Required: true,
			},
			"app_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAmplifyDomainCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Amplify::Domain", data, meta)
}

func resourceAmplifyDomainRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Amplify::Domain", data, meta)
}

func resourceAmplifyDomainUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Amplify::Domain", data, meta)
}

func resourceAmplifyDomainDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Amplify::Domain", data, meta)
}