// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html

package elasticache

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elastiCacheSecurityGroupIngressType string = "AWS::ElastiCache::SecurityGroupIngress"

func DatasourceElastiCacheSecurityGroupIngress() *schema.Resource {
	return &schema.Resource{
		Read: datasourceElastiCacheSecurityGroupIngressRead,
		
		Schema: map[string]*schema.Schema{
			"cache_security_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"ec2_security_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"ec2_security_group_owner_id": {
				Type: schema.TypeString,
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

func datasourceElastiCacheSecurityGroupIngressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elastiCacheSecurityGroupIngressType, DatasourceElastiCacheSecurityGroupIngress(), data, meta)
}
