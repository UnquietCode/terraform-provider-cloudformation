// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-template.html

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSESTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSESTemplateCreate,
		Read:   resourceSESTemplateRead,
		Update: resourceSESTemplateUpdate,
		Delete: resourceSESTemplateDelete,

		Schema: map[string]*schema.Schema{
			"template": {
				Type: schema.TypeList,
				Elem: propertyTemplateTemplate(),
				Optional: true,
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

func resourceSESTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SES::Template", ResourceSESTemplate(), data, meta)
}

func resourceSESTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SES::Template", ResourceSESTemplate(), data, meta)
}

func resourceSESTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SES::Template", ResourceSESTemplate(), data, meta)
}

func resourceSESTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SES::Template", data, meta)
}