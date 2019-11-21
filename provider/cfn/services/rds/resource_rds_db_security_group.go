// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html

package rds

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const rDSDBSecurityGroupType string = "AWS::RDS::DBSecurityGroup"

func ResourceRDSDBSecurityGroup() *schema.Resource {
	return &schema.Resource{
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

func resourceRDSDBSecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(rDSDBSecurityGroupType, ResourceRDSDBSecurityGroup(), data, meta)
}

func resourceRDSDBSecurityGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(rDSDBSecurityGroupType, ResourceRDSDBSecurityGroup(), data, meta)
}

func resourceRDSDBSecurityGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(rDSDBSecurityGroupType, ResourceRDSDBSecurityGroup(), data, meta)
}

func resourceRDSDBSecurityGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(rDSDBSecurityGroupType, data, meta)
}

func resourceRDSDBSecurityGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(rDSDBSecurityGroupType, data, meta)
}
