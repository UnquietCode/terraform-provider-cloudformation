// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGlueCrawler() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlueCrawlerCreate,
		Read:   resourceGlueCrawlerRead,
		Update: resourceGlueCrawlerUpdate,
		Delete: resourceGlueCrawlerDelete,

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
				Elem: propertySchemaChangePolicy(),
				Required: false,
				MaxItems: 1,
			},
			"configuration": {
				Type: schema.TypeString,
				Required: false,
			},
			"schedule": {
				Type: schema.TypeList,
				Elem: propertySchedule(),
				Required: false,
				MaxItems: 1,
			},
			"database_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"targets": {
				Type: schema.TypeList,
				Elem: propertyTargets(),
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

func resourceGlueCrawlerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::Crawler", data, meta)
}

func resourceGlueCrawlerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::Crawler", data, meta)
}

func resourceGlueCrawlerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::Crawler", data, meta)
}

func resourceGlueCrawlerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::Crawler", data, meta)
}