// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html

package route53

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const route53RecordSetGroupType string = "AWS::Route53::RecordSetGroup"

func ResourceRoute53RecordSetGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceRoute53RecordSetGroupRead,
		Create: resourceRoute53RecordSetGroupCreate,
		Update: resourceRoute53RecordSetGroupUpdate,
		Delete: resourceRoute53RecordSetGroupDelete,
		CustomizeDiff: resourceRoute53RecordSetGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"comment": {
				Type: schema.TypeString,
				Optional: true,
			},
			"hosted_zone_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"hosted_zone_name": {
				Type: schema.TypeString,
				Optional: true,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceRoute53RecordSetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(route53RecordSetGroupType, ResourceRoute53RecordSetGroup(), data, meta)
}

func resourceRoute53RecordSetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(route53RecordSetGroupType, ResourceRoute53RecordSetGroup(), data, meta)
}

func resourceRoute53RecordSetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(route53RecordSetGroupType, ResourceRoute53RecordSetGroup(), data, meta)
}

func resourceRoute53RecordSetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(route53RecordSetGroupType, data, meta)
}

func resourceRoute53RecordSetGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(route53RecordSetGroupType, data, meta)
}
