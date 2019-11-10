// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppSyncResolver() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppSyncResolverCreate,
		Read:   resourceAppSyncResolverRead,
		Update: resourceAppSyncResolverUpdate,
		Delete: resourceAppSyncResolverDelete,

		Schema: map[string]*schema.Schema{
			"type_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resolver_arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"field_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"response_mapping_template_s3_location": {
				Type: schema.TypeString,
				Optional: true,
			},
			"pipeline_config": {
				Type: schema.TypeList,
				Elem: propertyResolverPipelineConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"data_source_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"request_mapping_template": {
				Type: schema.TypeString,
				Optional: true,
			},
			"response_mapping_template": {
				Type: schema.TypeString,
				Optional: true,
			},
			"kind": {
				Type: schema.TypeString,
				Optional: true,
			},
			"request_mapping_template_s3_location": {
				Type: schema.TypeString,
				Optional: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAppSyncResolverCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppSync::Resolver", ResourceAppSyncResolver(), data, meta)
}

func resourceAppSyncResolverRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppSync::Resolver", ResourceAppSyncResolver(), data, meta)
}

func resourceAppSyncResolverUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppSync::Resolver", ResourceAppSyncResolver(), data, meta)
}

func resourceAppSyncResolverDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppSync::Resolver", data, meta)
}