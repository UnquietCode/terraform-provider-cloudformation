// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceStage() *schema.Resource {
	return &schema.Resource{
		Create: resourceStageCreate,
		Read:   resourceStageRead,
		Update: resourceStageUpdate,
		Delete: resourceStageDelete,

		Schema: map[string]*schema.Schema{
			"access_log_setting": {
				Type: schema.TypeList,
				Elem: propertyStageAccessLogSetting(),
				Required: false,
				MaxItems: 1,
			},
			"cache_cluster_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"cache_cluster_size": {
				Type: schema.TypeString,
				Required: false,
			},
			"canary_setting": {
				Type: schema.TypeList,
				Elem: propertyStageCanarySetting(),
				Required: false,
				MaxItems: 1,
			},
			"client_certificate_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"deployment_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"documentation_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"method_settings": {
				Type: schema.TypeSet,
				Elem: propertyStageMethodSetting(),
				Required: false,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"stage_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"tracing_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"variables": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}

func resourceStageCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::Stage", data, meta)
}

func resourceStageRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::Stage", data, meta)
}

func resourceStageUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::Stage", data, meta)
}

func resourceStageDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::Stage", data, meta)
}