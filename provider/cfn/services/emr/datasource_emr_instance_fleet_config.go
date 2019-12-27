// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html

package emr

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eMRInstanceFleetConfigType string = "AWS::EMR::InstanceFleetConfig"

func DatasourceEMRInstanceFleetConfig() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEMRInstanceFleetConfigRead,
		
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"instance_fleet_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"instance_type_configs": {
				Type: schema.TypeSet,
				Elem: propertyInstanceFleetConfigInstanceTypeConfig(),
				Optional: true,
			},
			"launch_specifications": {
				Type: schema.TypeList,
				Elem: propertyInstanceFleetConfigInstanceFleetProvisioningSpecifications(),
				Optional: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"target_on_demand_capacity": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"target_spot_capacity": {
				Type: schema.TypeInt,
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

func datasourceEMRInstanceFleetConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eMRInstanceFleetConfigType, DatasourceEMRInstanceFleetConfig(), data, meta)
}
