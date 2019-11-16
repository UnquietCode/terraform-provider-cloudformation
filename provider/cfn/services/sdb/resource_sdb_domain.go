// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simpledb.html

package sdb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sDBDomainType string = "AWS::SDB::Domain"

var sDBDomainProperties map[string]string = map[string]string{
	"description": "Description",
}

func ResourceSDBDomain() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSDBDomainExists,
		Read: resourceSDBDomainRead,
		Create: resourceSDBDomainCreate,
		Update: resourceSDBDomainUpdate,
		Delete: resourceSDBDomainDelete,
		CustomizeDiff: resourceSDBDomainCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
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

func resourceSDBDomainExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSDBDomainRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sDBDomainType, ResourceSDBDomain(), data, meta)
}

func resourceSDBDomainCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sDBDomainType, ResourceSDBDomain(), data, sDBDomainProperties, meta)
}

func resourceSDBDomainUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sDBDomainType, ResourceSDBDomain(), data, sDBDomainProperties, meta)
}

func resourceSDBDomainDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sDBDomainType, data, meta)
}

func resourceSDBDomainCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sDBDomainType, data, meta)
}
