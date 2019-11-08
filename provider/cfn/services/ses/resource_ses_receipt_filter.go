// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSESReceiptFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceSESReceiptFilterCreate,
		Read:   resourceSESReceiptFilterRead,
		Update: resourceSESReceiptFilterUpdate,
		Delete: resourceSESReceiptFilterDelete,

		Schema: map[string]*schema.Schema{
			"filter": {
				Type: schema.TypeList,
				Elem: propertyFilter(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceSESReceiptFilterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SES::ReceiptFilter", data, meta)
}

func resourceSESReceiptFilterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SES::ReceiptFilter", data, meta)
}

func resourceSESReceiptFilterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SES::ReceiptFilter", data, meta)
}

func resourceSESReceiptFilterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SES::ReceiptFilter", data, meta)
}