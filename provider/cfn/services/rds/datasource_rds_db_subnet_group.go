// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnet-group.html

package rds

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const rDSDBSubnetGroupType string = "AWS::RDS::DBSubnetGroup"

func DatasourceRDSDBSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Read: datasourceRDSDBSubnetGroupRead,
		
		Schema: map[string]*schema.Schema{
			"db_subnet_group_description": {
				Type: schema.TypeString,
				Required: true,
			},
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"subnet_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
				Set: schema.HashString,
			},
			"tags": misc.PropertyTags(),
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

func datasourceRDSDBSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(rDSDBSubnetGroupType, DatasourceRDSDBSubnetGroup(), data, meta)
}
