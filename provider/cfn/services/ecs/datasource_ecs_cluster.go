// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceECSCluster() *schema.Resource {
	return &schema.Resource{
		Read: datasourceECSClusterRead,
		
		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cluster_settings": {
				Type: schema.TypeSet,
				Elem: propertyClusterClusterSetting(),
				Optional: true,
			},
			"tags": misc.PropertyTags(),
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

func datasourceECSClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eCSClusterType, DatasourceECSCluster(), data, meta)
}
