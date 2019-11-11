// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html

package redshift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRedshiftClusterSecurityGroupIngress() *schema.Resource {
	return &schema.Resource{
		Create: resourceRedshiftClusterSecurityGroupIngressCreate,
		Read:   resourceRedshiftClusterSecurityGroupIngressRead,
		Delete: resourceRedshiftClusterSecurityGroupIngressDelete,

		Schema: map[string]*schema.Schema{
			"cidrip": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"cluster_security_group_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ec2_security_group_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ec2_security_group_owner_id": {
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

func resourceRedshiftClusterSecurityGroupIngressCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Redshift::ClusterSecurityGroupIngress", ResourceRedshiftClusterSecurityGroupIngress(), data, meta)
}

func resourceRedshiftClusterSecurityGroupIngressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Redshift::ClusterSecurityGroupIngress", ResourceRedshiftClusterSecurityGroupIngress(), data, meta)
}

func resourceRedshiftClusterSecurityGroupIngressUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Redshift::ClusterSecurityGroupIngress", ResourceRedshiftClusterSecurityGroupIngress(), data, meta)
}

func resourceRedshiftClusterSecurityGroupIngressDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Redshift::ClusterSecurityGroupIngress", data, meta)
}