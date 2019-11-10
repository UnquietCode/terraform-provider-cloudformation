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

func ResourceInstanceFleetConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstanceFleetConfigCreate,
		Read:   resourceInstanceFleetConfigRead,
		Update: resourceInstanceFleetConfigUpdate,
		Delete: resourceInstanceFleetConfigDelete,

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

func resourceInstanceFleetConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EMR::InstanceFleetConfig", data, meta)
}

func resourceInstanceFleetConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EMR::InstanceFleetConfig", data, meta)
}

func resourceInstanceFleetConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EMR::InstanceFleetConfig", data, meta)
}

func resourceInstanceFleetConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EMR::InstanceFleetConfig", data, meta)
}