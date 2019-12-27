// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceKinesisAnalyticsV2ApplicationOutput() *schema.Resource {
	return &schema.Resource{
		Read: datasourceKinesisAnalyticsV2ApplicationOutputRead,
		
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"output": {
				Type: schema.TypeList,
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
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceKinesisAnalyticsV2ApplicationOutputRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kinesisAnalyticsV2ApplicationOutputType, DatasourceKinesisAnalyticsV2ApplicationOutput(), data, meta)
}
