// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyMessage() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"json_body": {
				Type: schema.TypeString,
				Required: false,
			},
			"action": {
				Type: schema.TypeString,
				Required: false,
			},
			"media_url": {
				Type: schema.TypeString,
				Required: false,
			},
			"time_to_live": {
				Type: schema.TypeInt,
				Required: false,
			},
			"image_small_icon_url": {
				Type: schema.TypeString,
				Required: false,
			},
			"image_url": {
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
			"silent_push": {
				Type: schema.TypeBool,
				Required: false,
			},
			"body": {
				Type: schema.TypeString,
				Required: false,
			},
			"raw_content": {
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