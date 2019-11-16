// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eMRInstanceFleetConfigType string = "AWS::EMR::InstanceFleetConfig"

var eMRInstanceFleetConfigProperties map[string]string = map[string]string{
	"cluster_id": "ClusterId",
	"instance_fleet_type": "InstanceFleetType",
	"instance_type_configs": "InstanceTypeConfigs",
	"launch_specifications": "LaunchSpecifications",
	"name": "Name",
	"target_on_demand_capacity": "TargetOnDemandCapacity",
	"target_spot_capacity": "TargetSpotCapacity",
}

func ResourceEMRInstanceFleetConfig() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEMRInstanceFleetConfigExists,
		Read: resourceEMRInstanceFleetConfigRead,
		Create: resourceEMRInstanceFleetConfigCreate,
		Update: resourceEMRInstanceFleetConfigUpdate,
		Delete: resourceEMRInstanceFleetConfigDelete,
		CustomizeDiff: resourceEMRInstanceFleetConfigCustomizeDiff,
		
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
				Type: schema.TypeSet,
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
			},
		},
	}
}

func resourceEMRInstanceFleetConfigExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEMRInstanceFleetConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eMRInstanceFleetConfigType, ResourceEMRInstanceFleetConfig(), data, meta)
}

func resourceEMRInstanceFleetConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eMRInstanceFleetConfigType, ResourceEMRInstanceFleetConfig(), data, eMRInstanceFleetConfigProperties, meta)
}

func resourceEMRInstanceFleetConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eMRInstanceFleetConfigType, ResourceEMRInstanceFleetConfig(), data, eMRInstanceFleetConfigProperties, meta)
}

func resourceEMRInstanceFleetConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eMRInstanceFleetConfigType, data, meta)
}

func resourceEMRInstanceFleetConfigCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eMRInstanceFleetConfigType, data, meta)
}
