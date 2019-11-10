// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simpledb.html

package sdb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSDBDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceSDBDomainCreate,
		Read:   resourceSDBDomainRead,
		Update: resourceSDBDomainUpdate,
		Delete: resourceSDBDomainDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSDBDomainCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SDB::Domain", ResourceSDBDomain(), data, meta)
}

func resourceSDBDomainRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SDB::Domain", ResourceSDBDomain(), data, meta)
}

func resourceSDBDomainUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SDB::Domain", ResourceSDBDomain(), data, meta)
}

func resourceSDBDomainDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SDB::Domain", data, meta)
}