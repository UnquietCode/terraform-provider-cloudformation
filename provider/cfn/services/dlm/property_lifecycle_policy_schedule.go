// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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
			"tags_to_add": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"create_rule": {
				Type: schema.TypeList,
				Elem: propertyLifecyclePolicyCreateRule(),
				Optional: true,
				MaxItems: 1,
			},
			"variable_tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"fast_restore_rule": {
				Type: schema.TypeList,
				Elem: propertyLifecyclePolicyFastRestoreRule(),
				Optional: true,
				MaxItems: 1,
			},
			"retain_rule": {
				Type: schema.TypeList,
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
