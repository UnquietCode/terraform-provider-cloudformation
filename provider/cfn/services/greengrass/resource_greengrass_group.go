// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGreengrassGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGreengrassGroupCreate,
		Read:   resourceGreengrassGroupRead,
		Update: resourceGreengrassGroupUpdate,
		Delete: resourceGreengrassGroupDelete,

		Schema: map[string]*schema.Schema{
			"role_attached_at": {
				Type: schema.TypeString,
				Computed: true,
			},
			"latest_version_arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"the_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"initial_version": {
				Type: schema.TypeList,
				Elem: propertyGroupGroupVersion(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceGreengrassGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::Group", ResourceGreengrassGroup(), data, meta)
}

func resourceGreengrassGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::Group", ResourceGreengrassGroup(), data, meta)
}

func resourceGreengrassGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::Group", ResourceGreengrassGroup(), data, meta)
}

func resourceGreengrassGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::Group", data, meta)
}