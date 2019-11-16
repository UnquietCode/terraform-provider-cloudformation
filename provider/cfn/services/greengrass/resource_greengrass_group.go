// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassGroupType string = "AWS::Greengrass::Group"

var greengrassGroupProperties map[string]string = map[string]string{
	"initial_version": "InitialVersion",
	"role_arn": "RoleArn",
	"tags": "Tags",
	"name": "Name",
}

func ResourceGreengrassGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassGroupExists,
		Read: resourceGreengrassGroupRead,
		Create: resourceGreengrassGroupCreate,
		Update: resourceGreengrassGroupUpdate,
		Delete: resourceGreengrassGroupDelete,
		CustomizeDiff: resourceGreengrassGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeSet,
				Elem: propertyGroupGroupVersion(),
				Optional: true,
				MaxItems: 1,
			},
			"role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeMap,
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
			},
		},
	}
}

func resourceGreengrassGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassGroupType, ResourceGreengrassGroup(), data, meta)
}

func resourceGreengrassGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassGroupType, ResourceGreengrassGroup(), data, greengrassGroupProperties, meta)
}

func resourceGreengrassGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassGroupType, ResourceGreengrassGroup(), data, greengrassGroupProperties, meta)
}

func resourceGreengrassGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassGroupType, data, meta)
}

func resourceGreengrassGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassGroupType, data, meta)
}
