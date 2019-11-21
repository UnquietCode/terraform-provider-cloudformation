// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html

package ecs

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eCSClusterType string = "AWS::ECS::Cluster"

func ResourceECSCluster() *schema.Resource {
	return &schema.Resource{
		Read: resourceECSClusterRead,
		Create: resourceECSClusterCreate,
		Update: resourceECSClusterUpdate,
		Delete: resourceECSClusterDelete,
		CustomizeDiff: resourceECSClusterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceECSClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eCSClusterType, ResourceECSCluster(), data, meta)
}

func resourceECSClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eCSClusterType, ResourceECSCluster(), data, meta)
}

func resourceECSClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eCSClusterType, ResourceECSCluster(), data, meta)
}

func resourceECSClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eCSClusterType, data, meta)
}

func resourceECSClusterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eCSClusterType, data, meta)
}
