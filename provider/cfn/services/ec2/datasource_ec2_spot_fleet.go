// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2SpotFleetType string = "AWS::EC2::SpotFleet"

func DatasourceEC2SpotFleet() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2SpotFleetRead,
		
		Schema: map[string]*schema.Schema{
			"spot_fleet_request_config_data": {
				Type: schema.TypeList,
				Elem: propertySpotFleetSpotFleetRequestConfigData(),
				Required: true,
				MaxItems: 1,
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

func datasourceEC2SpotFleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2SpotFleetType, DatasourceEC2SpotFleet(), data, meta)
}
