// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-application.html

package kinesisanalytics

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const kinesisAnalyticsApplicationType string = "AWS::KinesisAnalytics::Application"

func DatasourceKinesisAnalyticsApplication() *schema.Resource {
	return &schema.Resource{
		Read: datasourceKinesisAnalyticsApplicationRead,
		
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"inputs": {
				Type: schema.TypeList,
				Elem: propertyApplicationInput(),
				Required: true,
			},
			"application_description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"application_code": {
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

func datasourceKinesisAnalyticsApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kinesisAnalyticsApplicationType, DatasourceKinesisAnalyticsApplication(), data, meta)
}
