// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html

package redshift

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const redshiftClusterSecurityGroupIngressType string = "AWS::Redshift::ClusterSecurityGroupIngress"

func DatasourceRedshiftClusterSecurityGroupIngress() *schema.Resource {
	return &schema.Resource{
		Read: datasourceRedshiftClusterSecurityGroupIngressRead,
		
		Schema: map[string]*schema.Schema{
			"cidrip": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cluster_security_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"ec2_security_group_name": {
				Type: schema.TypeString,
				Optional: true,
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

func datasourceRedshiftClusterSecurityGroupIngressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(redshiftClusterSecurityGroupIngressType, DatasourceRedshiftClusterSecurityGroupIngress(), data, meta)
}
