// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGreengrassGroupVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceGreengrassGroupVersionCreate,
		Read:   resourceGreengrassGroupVersionRead,
		Delete: resourceGreengrassGroupVersionDelete,

		Schema: map[string]*schema.Schema{
			"logger_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"device_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"function_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"core_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"resource_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"connector_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"subscription_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"group_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGreengrassGroupVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::GroupVersion", ResourceGreengrassGroupVersion(), data, meta)
}

func resourceGreengrassGroupVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::GroupVersion", ResourceGreengrassGroupVersion(), data, meta)
}

func resourceGreengrassGroupVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::GroupVersion", ResourceGreengrassGroupVersion(), data, meta)
}

func resourceGreengrassGroupVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::GroupVersion", data, meta)
}