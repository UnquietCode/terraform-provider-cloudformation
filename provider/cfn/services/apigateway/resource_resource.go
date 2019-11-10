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

func ResourceResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceResourceCreate,
		Read:   resourceResourceRead,
		Update: resourceResourceUpdate,
		Delete: resourceResourceDelete,

		Schema: map[string]*schema.Schema{
			"parent_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"path_part": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceResourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::Resource", data, meta)
}

func resourceResourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::Resource", data, meta)
}

func resourceResourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::Resource", data, meta)
}

func resourceResourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::Resource", data, meta)
}