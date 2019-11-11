// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html

package route53

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRoute53RecordSetGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoute53RecordSetGroupCreate,
		Read:   resourceRoute53RecordSetGroupRead,
		Update: resourceRoute53RecordSetGroupUpdate,
		Delete: resourceRoute53RecordSetGroupDelete,

		Schema: map[string]*schema.Schema{
			"comment": {
				Type: schema.TypeString,
				Optional: true,
			},
			"hosted_zone_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"hosted_zone_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"record_sets": {
				Type: schema.TypeSet,
				Elem: propertyRecordSetGroupRecordSet(),
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

func resourceRoute53RecordSetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Route53::RecordSetGroup", ResourceRoute53RecordSetGroup(), data, meta)
}

func resourceRoute53RecordSetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Route53::RecordSetGroup", ResourceRoute53RecordSetGroup(), data, meta)
}

func resourceRoute53RecordSetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Route53::RecordSetGroup", ResourceRoute53RecordSetGroup(), data, meta)
}

func resourceRoute53RecordSetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Route53::RecordSetGroup", data, meta)
}