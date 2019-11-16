// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html

package glue

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var triggerActionProperties map[string]string = map[string]string{
	"notification_property": "NotificationProperty",
	"crawler_name": "CrawlerName",
	"timeout": "Timeout",
	"job_name": "JobName",
	"arguments": "Arguments",
	"security_configuration": "SecurityConfiguration",
}

func propertyTriggerAction(extras...string) *schema.Resource {
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
			"notification_property": {
				Type: schema.TypeSet,
				Elem: propertyTriggerNotificationProperty(),
				Optional: true,
				MaxItems: 1,
			},
			"crawler_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"timeout": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"job_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"arguments": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"security_configuration": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
