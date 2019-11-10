// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCrawler() *schema.Resource {
	return &schema.Resource{
		Create: resourceCrawlerCreate,
		Read:   resourceCrawlerRead,
		Update: resourceCrawlerUpdate,
		Delete: resourceCrawlerDelete,

		Schema: map[string]*schema.Schema{
			"role": {
				Type: schema.TypeString,
				Required: true,
			},
			"classifiers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"schema_change_policy": {
				Type: schema.TypeList,
				Elem: propertyCrawlerSchemaChangePolicy(),
				Required: false,
				MaxItems: 1,
			},
			"configuration": {
				Type: schema.TypeString,
				Required: false,
			},
			"schedule": {
				Type: schema.TypeList,
				Elem: propertyCrawlerSchedule(),
				Required: false,
				MaxItems: 1,
			},
			"database_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"targets": {
				Type: schema.TypeList,
				Elem: propertyCrawlerTargets(),
				Required: true,
				MaxItems: 1,
			},
			"crawler_security_configuration": {
				Type: schema.TypeString,
				Required: false,
			},
			"table_prefix": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceCrawlerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::Crawler", data, meta)
}

func resourceCrawlerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::Crawler", data, meta)
}

func resourceCrawlerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::Crawler", data, meta)
}

func resourceCrawlerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::Crawler", data, meta)
}