// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-ledger.html

package qldb

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const qLDBLedgerType string = "AWS::QLDB::Ledger"

func DatasourceQLDBLedger() *schema.Resource {
	return &schema.Resource{
		Read: datasourceQLDBLedgerRead,
		
		Schema: map[string]*schema.Schema{
			"permissions_mode": {
				Type: schema.TypeString,
				Required: true,
			},
			"deletion_protection": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceQLDBLedgerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(qLDBLedgerType, DatasourceQLDBLedger(), data, meta)
}
