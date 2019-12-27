// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html

package appsync

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appSyncDataSourceType string = "AWS::AppSync::DataSource"

func DatasourceAppSyncDataSource() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAppSyncDataSourceRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceAppSyncDataSourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appSyncDataSourceType, DatasourceAppSyncDataSource(), data, meta)
}
