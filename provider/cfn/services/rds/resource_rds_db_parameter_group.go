// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html

package rds

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const rDSDBParameterGroupType string = "AWS::RDS::DBParameterGroup"

func ResourceRDSDBParameterGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceRDSDBParameterGroupRead,
		Create: resourceRDSDBParameterGroupCreate,
		Update: resourceRDSDBParameterGroupUpdate,
		Delete: resourceRDSDBParameterGroupDelete,
		CustomizeDiff: resourceRDSDBParameterGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"family": {
				Type: schema.TypeString,
				Required: true,
			},
			"parameters": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
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

func resourceRDSDBParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(rDSDBParameterGroupType, ResourceRDSDBParameterGroup(), data, meta)
}

func resourceRDSDBParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(rDSDBParameterGroupType, ResourceRDSDBParameterGroup(), data, meta)
}

func resourceRDSDBParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(rDSDBParameterGroupType, ResourceRDSDBParameterGroup(), data, meta)
}

func resourceRDSDBParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(rDSDBParameterGroupType, data, meta)
}

func resourceRDSDBParameterGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(rDSDBParameterGroupType, data, meta)
}
