// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-crawler.html

package glue

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueCrawlerType string = "AWS::Glue::Crawler"

func DatasourceGlueCrawler() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGlueCrawlerRead,
		
		Schema: map[string]*schema.Schema{
			"role": {
				Type: schema.TypeString,
				Required: true,
			},
			"classifiers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"schema_change_policy": {
				Type: schema.TypeList,
				Elem: propertyCrawlerSchemaChangePolicy(),
				Optional: true,
				MaxItems: 1,
			},
			"configuration": {
				Type: schema.TypeString,
				Optional: true,
			},
			"schedule": {
				Type: schema.TypeList,
				Elem: propertyCrawlerSchedule(),
				Optional: true,
				MaxItems: 1,
			},
			"database_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"targets": {
				Type: schema.TypeList,
				Elem: propertyCrawlerTargets(),
				Required: true,
				MaxItems: 1,
			},
			"crawler_security_configuration": {
				Type: schema.TypeString,
				Optional: true,
			},
			"table_prefix": {
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

func datasourceGlueCrawlerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueCrawlerType, DatasourceGlueCrawler(), data, meta)
}
