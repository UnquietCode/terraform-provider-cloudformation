// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceMLTransform() *schema.Resource {
	return &schema.Resource{
		Create: resourceMLTransformCreate,
		Read:   resourceMLTransformRead,
		Update: resourceMLTransformUpdate,
		Delete: resourceMLTransformDelete,

		Schema: map[string]*schema.Schema{
			"role": {
				Type: schema.TypeString,
				Required: true,
			},
			"max_retries": {
				Type: schema.TypeInt,
				Required: false,
			},
			"worker_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"timeout": {
				Type: schema.TypeInt,
				Required: false,
			},
			"transform_parameters": {
				Type: schema.TypeList,
				Elem: propertyMLTransformTransformParameters(),
				Required: true,
				MaxItems: 1,
			},
			"input_record_tables": {
				Type: schema.TypeList,
				Elem: propertyMLTransformInputRecordTables(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"number_of_workers": {
				Type: schema.TypeInt,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
			"max_capacity": {
				Type: schema.TypeFloat,
				Required: false,
			},
		},
	}
}

func resourceMLTransformCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::MLTransform", data, meta)
}

func resourceMLTransformRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::MLTransform", data, meta)
}

func resourceMLTransformUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::MLTransform", data, meta)
}

func resourceMLTransformDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::MLTransform", data, meta)
}