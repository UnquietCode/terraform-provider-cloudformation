// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html

package eks

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eKSNodegroupType string = "AWS::EKS::Nodegroup"

func DatasourceEKSNodegroup() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEKSNodegroupRead,
		
		Schema: map[string]*schema.Schema{
			"scaling_config": {
				Type: schema.TypeList,
				Elem: propertyNodegroupScalingConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"labels": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"release_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"nodegroup_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"subnets": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"node_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"ami_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"force_update_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"remote_access": {
				Type: schema.TypeList,
				Elem: propertyNodegroupRemoteAccess(),
				Optional: true,
				MaxItems: 1,
			},
			"disk_size": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"cluster_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"instance_types": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
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

func datasourceEKSNodegroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eKSNodegroupType, DatasourceEKSNodegroup(), data, meta)
}
