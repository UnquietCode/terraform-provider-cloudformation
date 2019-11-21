// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simpledb.html

package sdb

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sDBDomainType string = "AWS::SDB::Domain"

func ResourceSDBDomain() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceSDBDomainRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sDBDomainType, ResourceSDBDomain(), data, meta)
}

func resourceSDBDomainCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sDBDomainType, ResourceSDBDomain(), data, meta)
}

func resourceSDBDomainUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sDBDomainType, ResourceSDBDomain(), data, meta)
}

func resourceSDBDomainDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sDBDomainType, data, meta)
}

func resourceSDBDomainCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sDBDomainType, data, meta)
}
