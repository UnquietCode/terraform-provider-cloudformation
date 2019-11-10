// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceModel() *schema.Resource {
	return &schema.Resource{
		Create: resourceModelCreate,
		Read:   resourceModelRead,
		Update: resourceModelUpdate,
		Delete: resourceModelDelete,

		Schema: map[string]*schema.Schema{
			"content_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schema": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceModelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::Model", data, meta)
}

func resourceModelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::Model", data, meta)
}

func resourceModelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::Model", data, meta)
}

func resourceModelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::Model", data, meta)
}