// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appSyncDataSourceType string = "AWS::AppSync::DataSource"

var appSyncDataSourceProperties map[string]string = map[string]string{
	"type": "Type",
	"description": "Description",
	"service_role_arn": "ServiceRoleArn",
	"http_config": "HttpConfig",
	"relational_database_config": "RelationalDatabaseConfig",
	"lambda_config": "LambdaConfig",
	"api_id": "ApiId",
	"name": "Name",
	"dynamo_db_config": "DynamoDBConfig",
	"elasticsearch_config": "ElasticsearchConfig",
}

func ResourceAppSyncDataSource() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAppSyncDataSourceExists,
		Read: resourceAppSyncDataSourceRead,
		Create: resourceAppSyncDataSourceCreate,
		Update: resourceAppSyncDataSourceUpdate,
		Delete: resourceAppSyncDataSourceDelete,
		CustomizeDiff: resourceAppSyncDataSourceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"service_role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"http_config": {
				Type: schema.TypeList,
				Elem: propertyDataSourceHttpConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"relational_database_config": {
				Type: schema.TypeList,
				Elem: propertyDataSourceRelationalDatabaseConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"lambda_config": {
				Type: schema.TypeList,
				Elem: propertyDataSourceLambdaConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"dynamo_db_config": {
				Type: schema.TypeList,
				Elem: propertyDataSourceDynamoDBConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"elasticsearch_config": {
				Type: schema.TypeList,
				Elem: propertyDataSourceElasticsearchConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAppSyncDataSourceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAppSyncDataSourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appSyncDataSourceType, ResourceAppSyncDataSource(), data, meta)
}

func resourceAppSyncDataSourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appSyncDataSourceType, ResourceAppSyncDataSource(), data, appSyncDataSourceProperties, meta)
}

func resourceAppSyncDataSourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appSyncDataSourceType, ResourceAppSyncDataSource(), data, appSyncDataSourceProperties, meta)
}

func resourceAppSyncDataSourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appSyncDataSourceType, data, meta)
}

func resourceAppSyncDataSourceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appSyncDataSourceType, data, meta)
}
