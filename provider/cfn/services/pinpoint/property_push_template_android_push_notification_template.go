// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-pushtemplate-androidpushnotificationtemplate.html

package pinpoint

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var pushTemplateAndroidPushNotificationTemplateProperties map[string]string = map[string]string{
	"action": "Action",
	"image_url": "ImageUrl",
	"small_image_icon_url": "SmallImageIconUrl",
	"title": "Title",
	"image_icon_url": "ImageIconUrl",
	"sound": "Sound",
	"body": "Body",
	"url": "Url",
}

func propertyPushTemplateAndroidPushNotificationTemplate(extras...string) *schema.Resource {
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
			"action": {
				Type: schema.TypeString,
				Optional: true,
			},
			"image_url": {
				Type: schema.TypeString,
				Optional: true,
			},
			"small_image_icon_url": {
				Type: schema.TypeString,
				Optional: true,
			},
			"title": {
				Type: schema.TypeString,
				Optional: true,
			},
			"image_icon_url": {
				Type: schema.TypeString,
				Optional: true,
			},
			"sound": {
				Type: schema.TypeString,
				Optional: true,
			},
			"body": {
				Type: schema.TypeString,
				Optional: true,
			},
			"url": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
