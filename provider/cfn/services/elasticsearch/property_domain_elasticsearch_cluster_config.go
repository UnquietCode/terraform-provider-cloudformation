// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
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
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dedicated_master_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"dedicated_master_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"dedicated_master_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"instance_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"zone_awareness_config": {
				Type: schema.TypeList,
				Elem: propertyDomainZoneAwarenessConfig(),
				Required: false,
				MaxItems: 1,
			},
			"zone_awareness_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}