// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html

package elasticsearch

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDomainElasticsearchClusterConfig(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dedicated_master_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"dedicated_master_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"dedicated_master_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"zone_awareness_config": {
				Type: schema.TypeSet,
				Elem: propertyDomainZoneAwarenessConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"zone_awareness_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
		},
	}
}
