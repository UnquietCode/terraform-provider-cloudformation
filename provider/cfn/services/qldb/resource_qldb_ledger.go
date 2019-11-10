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

func ResourceQLDBLedger() *schema.Resource {
	return &schema.Resource{
		Create: resourceQLDBLedgerCreate,
		Read:   resourceQLDBLedgerRead,
		Update: resourceQLDBLedgerUpdate,
		Delete: resourceQLDBLedgerDelete,

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

func resourceQLDBLedgerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::QLDB::Ledger", data, meta)
}

func resourceQLDBLedgerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::QLDB::Ledger", data, meta)
}

func resourceQLDBLedgerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::QLDB::Ledger", data, meta)
}

func resourceQLDBLedgerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::QLDB::Ledger", data, meta)
}