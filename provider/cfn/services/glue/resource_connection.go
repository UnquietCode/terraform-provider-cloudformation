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

func ResourceConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceConnectionCreate,
		Read:   resourceConnectionRead,
		Update: resourceConnectionUpdate,
		Delete: resourceConnectionDelete,

		Schema: map[string]*schema.Schema{
			"connection_input": {
				Type: schema.TypeList,
				Elem: propertyConnectionConnectionInput(),
				Required: true,
				MaxItems: 1,
			},
			"catalog_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceConnectionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::Connection", data, meta)
}

func resourceConnectionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::Connection", data, meta)
}

func resourceConnectionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::Connection", data, meta)
}

func resourceConnectionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::Connection", data, meta)
}