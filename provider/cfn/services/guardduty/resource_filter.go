// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package guardduty

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceFilterCreate,
		Read:   resourceFilterRead,
		Update: resourceFilterUpdate,
		Delete: resourceFilterDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"detector_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"finding_criteria": {
				Type: schema.TypeList,
				Elem: propertyFilterFindingCriteria(),
				Required: true,
				MaxItems: 1,
			},
			"rank": {
				Type: schema.TypeInt,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceFilterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GuardDuty::Filter", data, meta)
}

func resourceFilterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GuardDuty::Filter", data, meta)
}

func resourceFilterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GuardDuty::Filter", data, meta)
}

func resourceFilterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GuardDuty::Filter", data, meta)
}