// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-ledger.html

package qldb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceQLDBLedger() *schema.Resource {
	return &schema.Resource{
		Exists: resourceQLDBLedgerExists,
		Read: resourceQLDBLedgerRead,
		Create: resourceQLDBLedgerCreate,
		Update: resourceQLDBLedgerUpdate,
		Delete: resourceQLDBLedgerDelete,
		CustomizeDiff: resourceQLDBLedgerCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"permissions_mode": {
				Type: schema.TypeString,
				Required: true,
			},
			"deletion_protection": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceQLDBLedgerExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceQLDBLedgerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::QLDB::Ledger", ResourceQLDBLedger(), data, meta)
}

func resourceQLDBLedgerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::QLDB::Ledger", ResourceQLDBLedger(), data, meta)
}

func resourceQLDBLedgerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::QLDB::Ledger", ResourceQLDBLedger(), data, meta)
}

func resourceQLDBLedgerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::QLDB::Ledger", data, meta)
}

func resourceQLDBLedgerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

