// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-inputsecuritygroup.html

package medialive

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const mediaLiveInputSecurityGroupType string = "AWS::MediaLive::InputSecurityGroup"

func ResourceMediaLiveInputSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceMediaLiveInputSecurityGroupRead,
		Create: resourceMediaLiveInputSecurityGroupCreate,
		Update: resourceMediaLiveInputSecurityGroupUpdate,
		Delete: resourceMediaLiveInputSecurityGroupDelete,
		CustomizeDiff: resourceMediaLiveInputSecurityGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"whitelist_rules": {
				Type: schema.TypeList,
				Elem: propertyInputSecurityGroupInputWhitelistRuleCidr(),
				Optional: true,
			},
			"tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
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

func resourceMediaLiveInputSecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(mediaLiveInputSecurityGroupType, ResourceMediaLiveInputSecurityGroup(), data, meta)
}

func resourceMediaLiveInputSecurityGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(mediaLiveInputSecurityGroupType, ResourceMediaLiveInputSecurityGroup(), data, meta)
}

func resourceMediaLiveInputSecurityGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(mediaLiveInputSecurityGroupType, ResourceMediaLiveInputSecurityGroup(), data, meta)
}

func resourceMediaLiveInputSecurityGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(mediaLiveInputSecurityGroupType, data, meta)
}

func resourceMediaLiveInputSecurityGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(mediaLiveInputSecurityGroupType, data, meta)
}
