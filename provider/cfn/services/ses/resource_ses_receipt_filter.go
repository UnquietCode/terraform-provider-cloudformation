// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptfilter.html

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSESReceiptFilter() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSESReceiptFilterExists,
		Read: resourceSESReceiptFilterRead,
		Create: resourceSESReceiptFilterCreate,
		Update: resourceSESReceiptFilterUpdate,
		Delete: resourceSESReceiptFilterDelete,
		CustomizeDiff: resourceSESReceiptFilterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"filter": {
				Type: schema.TypeList,
				Elem: propertyReceiptFilterFilter(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSESReceiptFilterExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSESReceiptFilterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SES::ReceiptFilter", ResourceSESReceiptFilter(), data, meta)
}

func resourceSESReceiptFilterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SES::ReceiptFilter", ResourceSESReceiptFilter(), data, meta)
}

func resourceSESReceiptFilterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SES::ReceiptFilter", ResourceSESReceiptFilter(), data, meta)
}

func resourceSESReceiptFilterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SES::ReceiptFilter", data, meta)
}

func resourceSESReceiptFilterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

