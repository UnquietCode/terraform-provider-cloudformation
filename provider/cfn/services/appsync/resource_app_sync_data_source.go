// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppSyncDataSource() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppSyncDataSourceCreate,
		Read:   resourceAppSyncDataSourceRead,
		Update: resourceAppSyncDataSourceUpdate,
		Delete: resourceAppSyncDataSourceDelete,

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
				Elem: propertyHttpConfig(),
				Required: false,
				MaxItems: 1,
			},
			"relational_database_config": {
				Type: schema.TypeList,
				Elem: propertyRelationalDatabaseConfig(),
				Required: false,
				MaxItems: 1,
			},
			"lambda_config": {
				Type: schema.TypeList,
				Elem: propertyLambdaConfig(),
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
				Elem: propertyDynamoDBConfig(),
				Required: false,
				MaxItems: 1,
			},
			"elasticsearch_config": {
				Type: schema.TypeList,
				Elem: propertyElasticsearchConfig(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceAppSyncDataSourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppSync::DataSource", data, meta)
}

func resourceAppSyncDataSourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppSync::DataSource", data, meta)
}

func resourceAppSyncDataSourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppSync::DataSource", data, meta)
}

func resourceAppSyncDataSourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppSync::DataSource", data, meta)
}