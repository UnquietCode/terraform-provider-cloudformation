// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html

package emr

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eMRClusterType string = "AWS::EMR::Cluster"

func DatasourceEMRCluster() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEMRClusterRead,
		
		Schema: map[string]*schema.Schema{
			"additional_info": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
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
				Type: schema.TypeList,
				Elem: propertyClusterJobFlowInstancesConfig(),
				Required: true,
				MaxItems: 1,
			},
			"job_flow_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"kerberos_attributes": {
				Type: schema.TypeList,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceEMRClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eMRClusterType, DatasourceEMRCluster(), data, meta)
}
