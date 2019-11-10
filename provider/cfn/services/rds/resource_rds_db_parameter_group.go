// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRDSDBParameterGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceRDSDBParameterGroupCreate,
		Read:   resourceRDSDBParameterGroupRead,
		Update: resourceRDSDBParameterGroupUpdate,
		Delete: resourceRDSDBParameterGroupDelete,

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
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
		},
	}
}

func resourceRDSDBParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RDS::DBParameterGroup", ResourceRDSDBParameterGroup(), data, meta)
}

func resourceRDSDBParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RDS::DBParameterGroup", ResourceRDSDBParameterGroup(), data, meta)
}

func resourceRDSDBParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RDS::DBParameterGroup", ResourceRDSDBParameterGroup(), data, meta)
}

func resourceRDSDBParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RDS::DBParameterGroup", data, meta)
}