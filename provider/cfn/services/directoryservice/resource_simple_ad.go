// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package directoryservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSimpleAD() *schema.Resource {
	return &schema.Resource{
		Create: resourceSimpleADCreate,
		Read:   resourceSimpleADRead,
		Update: resourceSimpleADUpdate,
		Delete: resourceSimpleADDelete,

		Schema: map[string]*schema.Schema{
			"create_alias": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"enable_sso": {
				Type: schema.TypeBool,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"short_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"size": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vpc_settings": {
				Type: schema.TypeList,
				Elem: propertySimpleADVpcSettings(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceSimpleADCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DirectoryService::SimpleAD", data, meta)
}

func resourceSimpleADRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DirectoryService::SimpleAD", data, meta)
}

func resourceSimpleADUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DirectoryService::SimpleAD", data, meta)
}

func resourceSimpleADDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DirectoryService::SimpleAD", data, meta)
}