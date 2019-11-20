// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-crawler.html

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueCrawlerType string = "AWS::Glue::Crawler"

func ResourceGlueCrawler() *schema.Resource {
	return &schema.Resource{
		Read: resourceGlueCrawlerRead,
		Create: resourceGlueCrawlerCreate,
		Update: resourceGlueCrawlerUpdate,
		Delete: resourceGlueCrawlerDelete,
		CustomizeDiff: resourceGlueCrawlerCustomizeDiff,
		
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
				Type: schema.TypeSet,
				Elem: propertyCrawlerSchemaChangePolicy(),
				Optional: true,
				MaxItems: 1,
			},
			"configuration": {
				Type: schema.TypeString,
				Optional: true,
			},
			"schedule": {
				Type: schema.TypeSet,
				Elem: propertyCrawlerSchedule(),
				Optional: true,
				MaxItems: 1,
			},
			"database_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"targets": {
				Type: schema.TypeSet,
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
				Type: schema.TypeMap,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGlueCrawlerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueCrawlerType, ResourceGlueCrawler(), data, meta)
}

func resourceGlueCrawlerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(glueCrawlerType, ResourceGlueCrawler(), data, meta)
}

func resourceGlueCrawlerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(glueCrawlerType, ResourceGlueCrawler(), data, meta)
}

func resourceGlueCrawlerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(glueCrawlerType, data, meta)
}

func resourceGlueCrawlerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(glueCrawlerType, data, meta)
}
