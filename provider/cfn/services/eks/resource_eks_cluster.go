// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html

package eks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eKSClusterType string = "AWS::EKS::Cluster"

var eKSClusterProperties map[string]string = map[string]string{
	"version": "Version",
	"role_arn": "RoleArn",
	"resources_vpc_config": "ResourcesVpcConfig",
	"name": "Name",
}

func ResourceEKSCluster() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEKSClusterExists,
		Read: resourceEKSClusterRead,
		Create: resourceEKSClusterCreate,
		Update: resourceEKSClusterUpdate,
		Delete: resourceEKSClusterDelete,
		CustomizeDiff: resourceEKSClusterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"resources_vpc_config": {
				Type: schema.TypeList,
				Elem: propertyClusterResourcesVpcConfig(),
				Required: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
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

func resourceEKSClusterExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEKSClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eKSClusterType, ResourceEKSCluster(), data, meta)
}

func resourceEKSClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eKSClusterType, ResourceEKSCluster(), data, eKSClusterProperties, meta)
}

func resourceEKSClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eKSClusterType, ResourceEKSCluster(), data, eKSClusterProperties, meta)
}

func resourceEKSClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eKSClusterType, data, meta)
}

func resourceEKSClusterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eKSClusterType, data, meta)
}
