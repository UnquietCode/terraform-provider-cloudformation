// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html

package cloudformation

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloudFormationStackType string = "AWS::CloudFormation::Stack"

func ResourceCloudFormationStack() *schema.Resource {
	return &schema.Resource{
		Read: resourceCloudFormationStackRead,
		Create: resourceCloudFormationStackCreate,
		Update: resourceCloudFormationStackUpdate,
		Delete: resourceCloudFormationStackDelete,
		CustomizeDiff: resourceCloudFormationStackCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"notification_ar_ns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"parameters": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"template_url": {
				Type: schema.TypeString,
				Required: true,
			},
			"timeout_in_minutes": {
				Type: schema.TypeInt,
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

func resourceCloudFormationStackRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloudFormationStackType, ResourceCloudFormationStack(), data, meta)
}

func resourceCloudFormationStackCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cloudFormationStackType, ResourceCloudFormationStack(), data, meta)
}

func resourceCloudFormationStackUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cloudFormationStackType, ResourceCloudFormationStack(), data, meta)
}

func resourceCloudFormationStackDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cloudFormationStackType, data, meta)
}

func resourceCloudFormationStackCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cloudFormationStackType, data, meta)
}
