// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDataSource() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataSourceCreate,
		Read:   resourceDataSourceRead,
		Update: resourceDataSourceUpdate,
		Delete: resourceDataSourceDelete,

		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"service_role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"http_config": {
				Type: schema.TypeList,
				Elem: propertyDataSourceHttpConfig(),
				Required: false,
				MaxItems: 1,
			},
			"relational_database_config": {
				Type: schema.TypeList,
				Elem: propertyDataSourceRelationalDatabaseConfig(),
				Required: false,
				MaxItems: 1,
			},
			"lambda_config": {
				Type: schema.TypeList,
				Elem: propertyDataSourceLambdaConfig(),
				Required: false,
				MaxItems: 1,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dynamo_db_config": {
				Type: schema.TypeList,
				Elem: propertyDataSourceDynamoDBConfig(),
				Required: false,
				MaxItems: 1,
			},
			"elasticsearch_config": {
				Type: schema.TypeList,
				Elem: propertyDataSourceElasticsearchConfig(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceDataSourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppSync::DataSource", data, meta)
}

func resourceDataSourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppSync::DataSource", data, meta)
}

func resourceDataSourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppSync::DataSource", data, meta)
}

func resourceDataSourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppSync::DataSource", data, meta)
}