// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html

package eks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEKSCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceEKSClusterCreate,
		Read:   resourceEKSClusterRead,
		Update: resourceEKSClusterUpdate,
		Delete: resourceEKSClusterDelete,

		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type: schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"certificate_authority_data": {
				Type: schema.TypeString,
				Computed: true,
			},
			"version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resources_vpc_config": {
				Type: schema.TypeList,
				Elem: propertyClusterResourcesVpcConfig(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEKSClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EKS::Cluster", ResourceEKSCluster(), data, meta)
}

func resourceEKSClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EKS::Cluster", ResourceEKSCluster(), data, meta)
}

func resourceEKSClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EKS::Cluster", ResourceEKSCluster(), data, meta)
}

func resourceEKSClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EKS::Cluster", data, meta)
}