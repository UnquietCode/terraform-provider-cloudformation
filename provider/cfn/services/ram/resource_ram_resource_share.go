// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html

package ram

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const rAMResourceShareType string = "AWS::RAM::ResourceShare"

var rAMResourceShareProperties map[string]string = map[string]string{
	"principals": "Principals",
	"allow_external_principals": "AllowExternalPrincipals",
	"resource_arns": "ResourceArns",
	"tags": "Tags",
	"name": "Name",
}

func ResourceRAMResourceShare() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRAMResourceShareExists,
		Read: resourceRAMResourceShareRead,
		Create: resourceRAMResourceShareCreate,
		Update: resourceRAMResourceShareUpdate,
		Delete: resourceRAMResourceShareDelete,
		CustomizeDiff: resourceRAMResourceShareCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"principals": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"allow_external_principals": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"resource_arns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRAMResourceShareExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRAMResourceShareRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(rAMResourceShareType, ResourceRAMResourceShare(), data, meta)
}

func resourceRAMResourceShareCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(rAMResourceShareType, ResourceRAMResourceShare(), data, rAMResourceShareProperties, meta)
}

func resourceRAMResourceShareUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(rAMResourceShareType, ResourceRAMResourceShare(), data, rAMResourceShareProperties, meta)
}

func resourceRAMResourceShareDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(rAMResourceShareType, data, meta)
}

func resourceRAMResourceShareCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(rAMResourceShareType, data, meta)
}
