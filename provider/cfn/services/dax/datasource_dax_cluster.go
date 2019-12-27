// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html

package dax

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const dAXClusterType string = "AWS::DAX::Cluster"

func DatasourceDAXCluster() *schema.Resource {
	return &schema.Resource{
		Read: datasourceDAXClusterRead,
		
		Schema: map[string]*schema.Schema{
			"sse_specification": {
				Type: schema.TypeList,
				Elem: propertyClusterSSESpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"replication_factor": {
				Type: schema.TypeInt,
				Required: true,
			},
			"parameter_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"availability_zones": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"iam_role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"notification_topic_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"node_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"cluster_name": {
				Type: schema.TypeString,
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

func datasourceDAXClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dAXClusterType, DatasourceDAXCluster(), data, meta)
}
