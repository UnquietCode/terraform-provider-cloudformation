// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html

package greengrass

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassGroupVersionType string = "AWS::Greengrass::GroupVersion"

func ResourceGreengrassGroupVersion() *schema.Resource {
	return &schema.Resource{
		Read: resourceGreengrassGroupVersionRead,
		Create: resourceGreengrassGroupVersionCreate,
		Update: resourceGreengrassGroupVersionUpdate,
		Delete: resourceGreengrassGroupVersionDelete,
		CustomizeDiff: resourceGreengrassGroupVersionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"logger_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"device_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"function_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"core_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"resource_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"connector_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"subscription_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"group_id": {
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

func resourceGreengrassGroupVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassGroupVersionType, ResourceGreengrassGroupVersion(), data, meta)
}

func resourceGreengrassGroupVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassGroupVersionType, ResourceGreengrassGroupVersion(), data, meta)
}

func resourceGreengrassGroupVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassGroupVersionType, ResourceGreengrassGroupVersion(), data, meta)
}

func resourceGreengrassGroupVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassGroupVersionType, data, meta)
}

func resourceGreengrassGroupVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassGroupVersionType, data, meta)
}
