// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEMRCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceEMRClusterCreate,
		Read:   resourceEMRClusterRead,
		Update: resourceEMRClusterUpdate,
		Delete: resourceEMRClusterDelete,

		Schema: map[string]*schema.Schema{
			"master_public_dns": {
				Type: schema.TypeString,
				Computed: true,
			},
			"additional_info": {
				Type: schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
			"applications": {
				Type: schema.TypeSet,
				Elem: propertyClusterApplication(),
				Optional: true,
				ForceNew: true,
			},
			"auto_scaling_role": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"bootstrap_actions": {
				Type: schema.TypeSet,
				Elem: propertyClusterBootstrapActionConfig(),
				Optional: true,
				ForceNew: true,
			},
			"configurations": {
				Type: schema.TypeSet,
				Elem: propertyClusterConfiguration(),
				Optional: true,
				ForceNew: true,
			},
			"custom_ami_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ebs_root_volume_size": {
				Type: schema.TypeInt,
				Optional: true,
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
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"log_uri": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"release_label": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"scale_down_behavior": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"security_configuration": {
				Type: schema.TypeString,
				Optional: true,
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
				Optional: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"visible_to_all_users": {
				Type: schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceEMRClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EMR::Cluster", data, meta)
}

func resourceEMRClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EMR::Cluster", data, meta)
}

func resourceEMRClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EMR::Cluster", data, meta)
}

func resourceEMRClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EMR::Cluster", data, meta)
}