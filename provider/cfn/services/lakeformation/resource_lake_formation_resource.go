// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package lakeformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLakeFormationResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceLakeFormationResourceCreate,
		Read:   resourceLakeFormationResourceRead,
		Update: resourceLakeFormationResourceUpdate,
		Delete: resourceLakeFormationResourceDelete,

		Schema: map[string]*schema.Schema{
			"resource_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"use_service_linked_role": {
				Type: schema.TypeBool,
				Required: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceLakeFormationResourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::LakeFormation::Resource", data, meta)
}

func resourceLakeFormationResourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::LakeFormation::Resource", data, meta)
}

func resourceLakeFormationResourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::LakeFormation::Resource", data, meta)
}

func resourceLakeFormationResourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::LakeFormation::Resource", data, meta)
}