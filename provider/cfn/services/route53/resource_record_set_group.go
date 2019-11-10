// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package route53

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRecordSetGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceRecordSetGroupCreate,
		Read:   resourceRecordSetGroupRead,
		Update: resourceRecordSetGroupUpdate,
		Delete: resourceRecordSetGroupDelete,

		Schema: map[string]*schema.Schema{
			"comment": {
				Type: schema.TypeString,
				Required: false,
			},
			"hosted_zone_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"hosted_zone_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"record_sets": {
				Type: schema.TypeSet,
				Elem: propertyRecordSetGroupRecordSet(),
				Required: false,
			},
		},
	}
}

func resourceRecordSetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Route53::RecordSetGroup", data, meta)
}

func resourceRecordSetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Route53::RecordSetGroup", data, meta)
}

func resourceRecordSetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Route53::RecordSetGroup", data, meta)
}

func resourceRecordSetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Route53::RecordSetGroup", data, meta)
}