// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-ipset.html

package waf

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFIPSetType string = "AWS::WAF::IPSet"

func DatasourceWAFIPSet() *schema.Resource {
	return &schema.Resource{
		Read: datasourceWAFIPSetRead,
		
		Schema: map[string]*schema.Schema{
			"ip_set_descriptors": {
				Type: schema.TypeSet,
				Elem: propertyIPSetIPSetDescriptor(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
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

func datasourceWAFIPSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFIPSetType, DatasourceWAFIPSet(), data, meta)
}
