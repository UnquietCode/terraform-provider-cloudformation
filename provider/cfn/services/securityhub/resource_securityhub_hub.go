// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-hub.html

package securityhub

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSecurityHubHub() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSecurityHubHubExists,
		Read: resourceSecurityHubHubRead,
		Create: resourceSecurityHubHubCreate,
		Update: resourceSecurityHubHubUpdate,
		Delete: resourceSecurityHubHubDelete,
		CustomizeDiff: resourceSecurityHubHubCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSecurityHubHubExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSecurityHubHubRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SecurityHub::Hub", ResourceSecurityHubHub(), data, meta)
}

func resourceSecurityHubHubCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SecurityHub::Hub", ResourceSecurityHubHub(), data, meta)
}

func resourceSecurityHubHubUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SecurityHub::Hub", ResourceSecurityHubHub(), data, meta)
}

func resourceSecurityHubHubDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SecurityHub::Hub", data, meta)
}

func resourceSecurityHubHubCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::SecurityHub::Hub", data, meta)
}

