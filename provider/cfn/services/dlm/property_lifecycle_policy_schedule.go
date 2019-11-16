// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-schedule.html

package dlm

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
)

var lifecyclePolicyScheduleProperties map[string]string = map[string]string{
	"tags_to_add": "TagsToAdd",
	"create_rule": "CreateRule",
	"variable_tags": "VariableTags",
	"fast_restore_rule": "FastRestoreRule",
	"retain_rule": "RetainRule",
	"name": "Name",
	"copy_tags": "CopyTags",
}

func propertyLifecyclePolicySchedule(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tags_to_add": misc.PropertyTags(),
			"create_rule": {
				Type: schema.TypeSet,
				Elem: propertyLifecyclePolicyCreateRule(),
				Optional: true,
				MaxItems: 1,
			},
			"variable_tags": misc.PropertyTags(),
			"fast_restore_rule": {
				Type: schema.TypeSet,
				Elem: propertyLifecyclePolicyFastRestoreRule(),
				Optional: true,
				MaxItems: 1,
			},
			"retain_rule": {
				Type: schema.TypeSet,
				Elem: propertyLifecyclePolicyRetainRule(),
				Optional: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"copy_tags": {
				Type: schema.TypeBool,
				Optional: true,
			},
		},
	}
}
