// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package qldb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLedger() *schema.Resource {
	return &schema.Resource{
		Create: resourceLedgerCreate,
		Read:   resourceLedgerRead,
		Update: resourceLedgerUpdate,
		Delete: resourceLedgerDelete,

		Schema: map[string]*schema.Schema{
			"permissions_mode": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"deletion_protection": {
				Type: schema.TypeBool,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceLedgerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::QLDB::Ledger", data, meta)
}

func resourceLedgerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::QLDB::Ledger", data, meta)
}

func resourceLedgerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::QLDB::Ledger", data, meta)
}

func resourceLedgerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::QLDB::Ledger", data, meta)
}