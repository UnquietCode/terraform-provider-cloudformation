// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html

package appsync

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appSyncResolverType string = "AWS::AppSync::Resolver"

func ResourceAppSyncResolver() *schema.Resource {
	return &schema.Resource{
		Read: resourceAppSyncResolverRead,
		Create: resourceAppSyncResolverCreate,
		Update: resourceAppSyncResolverUpdate,
		Delete: resourceAppSyncResolverDelete,
		CustomizeDiff: resourceAppSyncResolverCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"response_mapping_template_s3_location": {
				Type: schema.TypeString,
				Optional: true,
			},
			"type_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"pipeline_config": {
				Type: schema.TypeSet,
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
			},
			"field_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceAppSyncResolverRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appSyncResolverType, ResourceAppSyncResolver(), data, meta)
}

func resourceAppSyncResolverCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appSyncResolverType, ResourceAppSyncResolver(), data, meta)
}

func resourceAppSyncResolverUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appSyncResolverType, ResourceAppSyncResolver(), data, meta)
}

func resourceAppSyncResolverDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appSyncResolverType, data, meta)
}

func resourceAppSyncResolverCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appSyncResolverType, data, meta)
}
