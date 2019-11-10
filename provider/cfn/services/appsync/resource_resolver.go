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

func ResourceResolver() *schema.Resource {
	return &schema.Resource{
		Create: resourceResolverCreate,
		Read:   resourceResolverRead,
		Update: resourceResolverUpdate,
		Delete: resourceResolverDelete,

		Schema: map[string]*schema.Schema{
			"response_mapping_template_s3_location": {
				Type: schema.TypeString,
				Required: false,
			},
			"type_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pipeline_config": {
				Type: schema.TypeList,
				Elem: propertyResolverPipelineConfig(),
				Required: false,
				MaxItems: 1,
			},
			"data_source_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"request_mapping_template": {
				Type: schema.TypeString,
				Required: false,
			},
			"response_mapping_template": {
				Type: schema.TypeString,
				Required: false,
			},
			"kind": {
				Type: schema.TypeString,
				Required: false,
			},
			"request_mapping_template_s3_location": {
				Type: schema.TypeString,
				Required: false,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"field_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceResolverCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppSync::Resolver", data, meta)
}

func resourceResolverRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppSync::Resolver", data, meta)
}

func resourceResolverUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppSync::Resolver", data, meta)
}

func resourceResolverDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppSync::Resolver", data, meta)
}