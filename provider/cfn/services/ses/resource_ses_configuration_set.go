// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationset.html

package ses

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sESConfigurationSetType string = "AWS::SES::ConfigurationSet"

func ResourceSESConfigurationSet() *schema.Resource {
	return &schema.Resource{
		Read: resourceSESConfigurationSetRead,
		Create: resourceSESConfigurationSetCreate,
		Update: resourceSESConfigurationSetUpdate,
		Delete: resourceSESConfigurationSetDelete,
		CustomizeDiff: resourceSESConfigurationSetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceSESConfigurationSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sESConfigurationSetType, ResourceSESConfigurationSet(), data, meta)
}

func resourceSESConfigurationSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sESConfigurationSetType, ResourceSESConfigurationSet(), data, meta)
}

func resourceSESConfigurationSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sESConfigurationSetType, ResourceSESConfigurationSet(), data, meta)
}

func resourceSESConfigurationSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sESConfigurationSetType, data, meta)
}

func resourceSESConfigurationSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sESConfigurationSetType, data, meta)
}
