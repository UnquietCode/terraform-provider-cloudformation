// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html

package rds

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const rDSOptionGroupType string = "AWS::RDS::OptionGroup"

func DatasourceRDSOptionGroup() *schema.Resource {
	return &schema.Resource{
		Read: datasourceRDSOptionGroupRead,
		
		Schema: map[string]*schema.Schema{
			"engine_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"major_engine_version": {
				Type: schema.TypeString,
				Required: true,
			},
			"option_configurations": {
				Type: schema.TypeList,
				Elem: propertyOptionGroupOptionConfiguration(),
				Required: true,
			},
			"option_group_description": {
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
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceRDSOptionGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(rDSOptionGroupType, DatasourceRDSOptionGroup(), data, meta)
}
