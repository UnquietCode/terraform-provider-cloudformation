// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html

package mediaconvert

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const mediaConvertJobTemplateType string = "AWS::MediaConvert::JobTemplate"

func DatasourceMediaConvertJobTemplate() *schema.Resource {
	return &schema.Resource{
		Read: datasourceMediaConvertJobTemplateRead,
		
		Schema: map[string]*schema.Schema{
			"category": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"acceleration_settings": {
				Type: schema.TypeList,
				Elem: propertyJobTemplateAccelerationSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"priority": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"status_update_interval": {
				Type: schema.TypeString,
				Optional: true,
			},
			"settings_json": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"queue": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"name": {
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

func datasourceMediaConvertJobTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(mediaConvertJobTemplateType, DatasourceMediaConvertJobTemplate(), data, meta)
}
