// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html

package ecs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceECSCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceECSClusterCreate,
		Read:   resourceECSClusterRead,
		Update: resourceECSClusterUpdate,
		Delete: resourceECSClusterDelete,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"cluster_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceECSClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ECS::Cluster", ResourceECSCluster(), data, meta)
}

func resourceECSClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ECS::Cluster", ResourceECSCluster(), data, meta)
}

func resourceECSClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ECS::Cluster", ResourceECSCluster(), data, meta)
}

func resourceECSClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ECS::Cluster", data, meta)
}