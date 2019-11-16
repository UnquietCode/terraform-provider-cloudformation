// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassGroupVersionType string = "AWS::Greengrass::GroupVersion"

var greengrassGroupVersionProperties map[string]string = map[string]string{
	"logger_definition_version_arn": "LoggerDefinitionVersionArn",
	"device_definition_version_arn": "DeviceDefinitionVersionArn",
	"function_definition_version_arn": "FunctionDefinitionVersionArn",
	"core_definition_version_arn": "CoreDefinitionVersionArn",
	"resource_definition_version_arn": "ResourceDefinitionVersionArn",
	"connector_definition_version_arn": "ConnectorDefinitionVersionArn",
	"subscription_definition_version_arn": "SubscriptionDefinitionVersionArn",
	"group_id": "GroupId",
}

func ResourceGreengrassGroupVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassGroupVersionExists,
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
			},
		},
	}
}

func resourceGreengrassGroupVersionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassGroupVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassGroupVersionType, ResourceGreengrassGroupVersion(), data, meta)
}

func resourceGreengrassGroupVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassGroupVersionType, ResourceGreengrassGroupVersion(), data, greengrassGroupVersionProperties, meta)
}

func resourceGreengrassGroupVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassGroupVersionType, ResourceGreengrassGroupVersion(), data, greengrassGroupVersionProperties, meta)
}

func resourceGreengrassGroupVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassGroupVersionType, data, meta)
}

func resourceGreengrassGroupVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassGroupVersionType, data, meta)
}
