// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationoutput.html

package kinesisanalyticsv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const kinesisAnalyticsV2ApplicationOutputType string = "AWS::KinesisAnalyticsV2::ApplicationOutput"

func ResourceKinesisAnalyticsV2ApplicationOutput() *schema.Resource {
	return &schema.Resource{
		Read: resourceKinesisAnalyticsV2ApplicationOutputRead,
		Create: resourceKinesisAnalyticsV2ApplicationOutputCreate,
		Update: resourceKinesisAnalyticsV2ApplicationOutputUpdate,
		Delete: resourceKinesisAnalyticsV2ApplicationOutputDelete,
		CustomizeDiff: resourceKinesisAnalyticsV2ApplicationOutputCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"output": {
				Type: schema.TypeSet,
				Elem: propertyApplicationOutputOutput(),
				Required: true,
				MaxItems: 1,
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

func resourceKinesisAnalyticsV2ApplicationOutputRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kinesisAnalyticsV2ApplicationOutputType, ResourceKinesisAnalyticsV2ApplicationOutput(), data, meta)
}

func resourceKinesisAnalyticsV2ApplicationOutputCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(kinesisAnalyticsV2ApplicationOutputType, ResourceKinesisAnalyticsV2ApplicationOutput(), data, meta)
}

func resourceKinesisAnalyticsV2ApplicationOutputUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(kinesisAnalyticsV2ApplicationOutputType, ResourceKinesisAnalyticsV2ApplicationOutput(), data, meta)
}

func resourceKinesisAnalyticsV2ApplicationOutputDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(kinesisAnalyticsV2ApplicationOutputType, data, meta)
}

func resourceKinesisAnalyticsV2ApplicationOutputCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(kinesisAnalyticsV2ApplicationOutputType, data, meta)
}
