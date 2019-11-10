// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

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
				Required: false,
				ForceNew: true,
			},
			"launch_specifications": {
				Type: schema.TypeList,
				Elem: propertyInstanceFleetConfigInstanceFleetProvisioningSpecifications(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"target_on_demand_capacity": {
				Type: schema.TypeInt,
				Required: false,
			},
			"target_spot_capacity": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}

func resourceEMRInstanceFleetConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EMR::InstanceFleetConfig", data, meta)
}

func resourceEMRInstanceFleetConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EMR::InstanceFleetConfig", data, meta)
}

func resourceEMRInstanceFleetConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EMR::InstanceFleetConfig", data, meta)
}

func resourceEMRInstanceFleetConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EMR::InstanceFleetConfig", data, meta)
}