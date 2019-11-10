// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package pinpoint

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyPushTemplateAndroidPushNotificationTemplate(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString,
				Required: false,
			},
			"image_url": {
				Type: schema.TypeString,
				Required: false,
			},
			"small_image_icon_url": {
				Type: schema.TypeString,
				Required: false,
			},
			"title": {
				Type: schema.TypeString,
				Required: false,
			},
			"image_icon_url": {
				Type: schema.TypeString,
				Required: false,
			},
			"sound": {
				Type: schema.TypeString,
				Required: false,
			},
			"body": {
				Type: schema.TypeString,
				Required: false,
			},
			"url": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}