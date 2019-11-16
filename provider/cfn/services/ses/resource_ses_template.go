// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-template.html

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sESTemplateType string = "AWS::SES::Template"

func ResourceSESTemplate() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSESTemplateExists,
		Read: resourceSESTemplateRead,
		Create: resourceSESTemplateCreate,
		Update: resourceSESTemplateUpdate,
		Delete: resourceSESTemplateDelete,
		CustomizeDiff: resourceSESTemplateCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"template": {
				Type: schema.TypeSet,
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

func resourceSESTemplateExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSESTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sESTemplateType, ResourceSESTemplate(), data, meta)
}

func resourceSESTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sESTemplateType, ResourceSESTemplate(), data, meta)
}

func resourceSESTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sESTemplateType, ResourceSESTemplate(), data, meta)
}

func resourceSESTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sESTemplateType, data, meta)
}

func resourceSESTemplateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sESTemplateType, data, meta)
}
