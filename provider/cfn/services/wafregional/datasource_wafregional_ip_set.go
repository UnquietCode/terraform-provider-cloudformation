// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ipset.html

package wafregional

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFRegionalIPSetType string = "AWS::WAFRegional::IPSet"

func DatasourceWAFRegionalIPSet() *schema.Resource {
	return &schema.Resource{
		Read: datasourceWAFRegionalIPSetRead,
		
		Schema: map[string]*schema.Schema{
			"ip_set_descriptors": {
				Type: schema.TypeList,
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

func datasourceWAFRegionalIPSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFRegionalIPSetType, DatasourceWAFRegionalIPSet(), data, meta)
}
