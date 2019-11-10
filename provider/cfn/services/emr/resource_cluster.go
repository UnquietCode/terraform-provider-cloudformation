// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: resourceClusterDelete,

		Schema: map[string]*schema.Schema{
			"additional_info": {
				Type: schema.TypeMap,
				Required: false,
				ForceNew: true,
			},
			"applications": {
				Type: schema.TypeSet,
				Elem: propertyClusterApplication(),
				Required: false,
				ForceNew: true,
			},
			"auto_scaling_role": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"bootstrap_actions": {
				Type: schema.TypeSet,
				Elem: propertyClusterBootstrapActionConfig(),
				Required: false,
				ForceNew: true,
			},
			"configurations": {
				Type: schema.TypeSet,
				Elem: propertyClusterConfiguration(),
				Required: false,
				ForceNew: true,
			},
			"custom_ami_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"ebs_root_volume_size": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"instances": {
				Type: schema.TypeList,
				Elem: propertyClusterJobFlowInstancesConfig(),
				Required: true,
				MaxItems: 1,
			},
			"job_flow_role": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"kerberos_attributes": {
				Type: schema.TypeList,
				Elem: propertyClusterKerberosAttributes(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"log_uri": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"release_label": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"scale_down_behavior": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"security_configuration": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"service_role": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"steps": {
				Type: schema.TypeSet,
				Elem: propertyClusterStepConfig(),
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"visible_to_all_users": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}

func resourceClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EMR::Cluster", data, meta)
}

func resourceClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EMR::Cluster", data, meta)
}

func resourceClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EMR::Cluster", data, meta)
}

func resourceClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EMR::Cluster", data, meta)
}