// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEMRInstanceFleetConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceEMRInstanceFleetConfigCreate,
		Read:   resourceEMRInstanceFleetConfigRead,
		Update: resourceEMRInstanceFleetConfigUpdate,
		Delete: resourceEMRInstanceFleetConfigDelete,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_fleet_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_type_configs": {
				Type: schema.TypeSet,
				Elem: propertyInstanceFleetConfigInstanceTypeConfig(),
				Optional: true,
				ForceNew: true,
			},
			"launch_specifications": {
				Type: schema.TypeList,
				Elem: propertyInstanceFleetConfigInstanceFleetProvisioningSpecifications(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
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
			},
		},
	}
}

func resourceEMRInstanceFleetConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EMR::InstanceFleetConfig", ResourceEMRInstanceFleetConfig(), data, meta)
}

func resourceEMRInstanceFleetConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EMR::InstanceFleetConfig", ResourceEMRInstanceFleetConfig(), data, meta)
}

func resourceEMRInstanceFleetConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EMR::InstanceFleetConfig", ResourceEMRInstanceFleetConfig(), data, meta)
}

func resourceEMRInstanceFleetConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EMR::InstanceFleetConfig", data, meta)
}