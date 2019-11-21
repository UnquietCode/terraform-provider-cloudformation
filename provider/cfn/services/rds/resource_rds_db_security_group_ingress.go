// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html

package rds

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const rDSDBSecurityGroupIngressType string = "AWS::RDS::DBSecurityGroupIngress"

func ResourceRDSDBSecurityGroupIngress() *schema.Resource {
	return &schema.Resource{
		Read: resourceRDSDBSecurityGroupIngressRead,
		Create: resourceRDSDBSecurityGroupIngressCreate,
		Update: resourceRDSDBSecurityGroupIngressUpdate,
		Delete: resourceRDSDBSecurityGroupIngressDelete,
		CustomizeDiff: resourceRDSDBSecurityGroupIngressCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"cidrip": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_security_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"ec2_security_group_id": {
				Type: schema.TypeString,
				Optional: true,
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
		},
	}
}

func resourceRDSDBSecurityGroupIngressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(rDSDBSecurityGroupIngressType, ResourceRDSDBSecurityGroupIngress(), data, meta)
}

func resourceRDSDBSecurityGroupIngressCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(rDSDBSecurityGroupIngressType, ResourceRDSDBSecurityGroupIngress(), data, meta)
}

func resourceRDSDBSecurityGroupIngressUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(rDSDBSecurityGroupIngressType, ResourceRDSDBSecurityGroupIngress(), data, meta)
}

func resourceRDSDBSecurityGroupIngressDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(rDSDBSecurityGroupIngressType, data, meta)
}

func resourceRDSDBSecurityGroupIngressCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(rDSDBSecurityGroupIngressType, data, meta)
}
