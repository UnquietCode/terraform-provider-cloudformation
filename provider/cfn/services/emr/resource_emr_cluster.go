// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

const eMRClusterType string = "AWS::EMR::Cluster"

var eMRClusterProperties map[string]string = map[string]string{
	"additional_info": "AdditionalInfo",
	"applications": "Applications",
	"auto_scaling_role": "AutoScalingRole",
	"bootstrap_actions": "BootstrapActions",
	"configurations": "Configurations",
	"custom_ami_id": "CustomAmiId",
	"ebs_root_volume_size": "EbsRootVolumeSize",
	"instances": "Instances",
	"job_flow_role": "JobFlowRole",
	"kerberos_attributes": "KerberosAttributes",
	"log_uri": "LogUri",
	"name": "Name",
	"release_label": "ReleaseLabel",
	"scale_down_behavior": "ScaleDownBehavior",
	"security_configuration": "SecurityConfiguration",
	"service_role": "ServiceRole",
	"steps": "Steps",
	"tags": "Tags",
	"visible_to_all_users": "VisibleToAllUsers",
}

func ResourceEMRCluster() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEMRClusterExists,
		Read: resourceEMRClusterRead,
		Create: resourceEMRClusterCreate,
		Update: resourceEMRClusterUpdate,
		Delete: resourceEMRClusterDelete,
		CustomizeDiff: resourceEMRClusterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"additional_info": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"applications": {
				Type: schema.TypeSet,
				Elem: propertyClusterApplication(),
				Optional: true,
			},
			"auto_scaling_role": {
				Type: schema.TypeString,
				Optional: true,
			},
			"bootstrap_actions": {
				Type: schema.TypeSet,
				Elem: propertyClusterBootstrapActionConfig(),
				Optional: true,
			},
			"configurations": {
				Type: schema.TypeSet,
				Elem: propertyClusterConfiguration(),
				Optional: true,
			},
			"custom_ami_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ebs_root_volume_size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"instances": {
				Type: schema.TypeSet,
				Elem: propertyClusterJobFlowInstancesConfig(),
				Required: true,
				MaxItems: 1,
			},
			"job_flow_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"kerberos_attributes": {
				Type: schema.TypeSet,
				Elem: propertyClusterKerberosAttributes(),
				Optional: true,
				MaxItems: 1,
			},
			"log_uri": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"release_label": {
				Type: schema.TypeString,
				Optional: true,
			},
			"scale_down_behavior": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_configuration": {
				Type: schema.TypeString,
				Optional: true,
			},
			"service_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"steps": {
				Type: schema.TypeSet,
				Elem: propertyClusterStepConfig(),
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"visible_to_all_users": {
				Type: schema.TypeBool,
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

func resourceEMRClusterExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEMRClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eMRClusterType, ResourceEMRCluster(), data, meta)
}

func resourceEMRClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eMRClusterType, ResourceEMRCluster(), data, eMRClusterProperties, meta)
}

func resourceEMRClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eMRClusterType, ResourceEMRCluster(), data, eMRClusterProperties, meta)
}

func resourceEMRClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eMRClusterType, data, meta)
}

func resourceEMRClusterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eMRClusterType, data, meta)
}
