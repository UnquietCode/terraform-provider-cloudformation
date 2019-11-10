// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package athena

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAthenaNamedQuery() *schema.Resource {
	return &schema.Resource{
		Create: resourceAthenaNamedQueryCreate,
		Read:   resourceAthenaNamedQueryRead,
		Update: resourceAthenaNamedQueryUpdate,
		Delete: resourceAthenaNamedQueryDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"query_string": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"database": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceAthenaNamedQueryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Athena::NamedQuery", data, meta)
}

func resourceAthenaNamedQueryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Athena::NamedQuery", data, meta)
}

func resourceAthenaNamedQueryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Athena::NamedQuery", data, meta)
}

func resourceAthenaNamedQueryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Athena::NamedQuery", data, meta)
}