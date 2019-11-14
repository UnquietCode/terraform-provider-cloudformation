// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRDSDBSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRDSDBSecurityGroupExists,
		Read: resourceRDSDBSecurityGroupRead,
		Create: resourceRDSDBSecurityGroupCreate,
		Update: resourceRDSDBSecurityGroupUpdate,
		Delete: resourceRDSDBSecurityGroupDelete,
		CustomizeDiff: resourceRDSDBSecurityGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"db_security_group_ingress": {
				Type: schema.TypeSet,
				Elem: propertyDBSecurityGroupIngress(),
				Required: true,
			},
			"ec2_vpc_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"group_description": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRDSDBSecurityGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRDSDBSecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RDS::DBSecurityGroup", ResourceRDSDBSecurityGroup(), data, meta)
}

func resourceRDSDBSecurityGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RDS::DBSecurityGroup", ResourceRDSDBSecurityGroup(), data, meta)
}

func resourceRDSDBSecurityGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RDS::DBSecurityGroup", ResourceRDSDBSecurityGroup(), data, meta)
}

func resourceRDSDBSecurityGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RDS::DBSecurityGroup", data, meta)
}

func resourceRDSDBSecurityGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
