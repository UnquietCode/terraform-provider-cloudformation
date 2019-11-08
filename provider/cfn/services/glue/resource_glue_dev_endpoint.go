// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGlueDevEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlueDevEndpointCreate,
		Read:   resourceGlueDevEndpointRead,
		Update: resourceGlueDevEndpointUpdate,
		Delete: resourceGlueDevEndpointDelete,

		Schema: map[string]*schema.Schema{
			"extra_jars_s3_path": {
				Type: schema.TypeString,
				Required: false,
			},
			"public_key": {
				Type: schema.TypeString,
				Required: false,
			},
			"number_of_nodes": {
				Type: schema.TypeInt,
				Required: false,
			},
			"arguments": {
				Type: schema.TypeMap,
				Required: false,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"worker_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"endpoint_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"glue_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"extra_python_libs_s3_path": {
				Type: schema.TypeString,
				Required: false,
			},
			"security_configuration": {
				Type: schema.TypeString,
				Required: false,
			},
			"number_of_workers": {
				Type: schema.TypeInt,
				Required: false,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceGlueDevEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::DevEndpoint", data, meta)
}

func resourceGlueDevEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::DevEndpoint", data, meta)
}

func resourceGlueDevEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::DevEndpoint", data, meta)
}

func resourceGlueDevEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::DevEndpoint", data, meta)
}