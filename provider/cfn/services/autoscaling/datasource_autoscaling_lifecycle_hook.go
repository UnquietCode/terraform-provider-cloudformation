// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html

package autoscaling

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const autoScalingLifecycleHookType string = "AWS::AutoScaling::LifecycleHook"

func DatasourceAutoScalingLifecycleHook() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAutoScalingLifecycleHookRead,
		
		Schema: map[string]*schema.Schema{
			"auto_scaling_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"default_result": {
				Type: schema.TypeString,
				Optional: true,
			},
			"heartbeat_timeout": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"lifecycle_hook_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"lifecycle_transition": {
				Type: schema.TypeString,
				Required: true,
			},
			"notification_metadata": {
				Type: schema.TypeString,
				Optional: true,
			},
			"notification_target_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceAutoScalingLifecycleHookRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(autoScalingLifecycleHookType, DatasourceAutoScalingLifecycleHook(), data, meta)
}
