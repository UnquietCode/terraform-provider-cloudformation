// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-stacksetconstraint.html

package servicecatalog

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceCatalogStackSetConstraintType string = "AWS::ServiceCatalog::StackSetConstraint"

func DatasourceServiceCatalogStackSetConstraint() *schema.Resource {
	return &schema.Resource{
		Read: datasourceServiceCatalogStackSetConstraintRead,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"stack_instance_control": {
				Type: schema.TypeString,
				Required: true,
			},
			"accept_language": {
				Type: schema.TypeString,
				Optional: true,
			},
			"portfolio_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"product_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"region_list": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"admin_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"account_list": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"execution_role": {
				Type: schema.TypeString,
				Required: true,
			},
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

func datasourceServiceCatalogStackSetConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceCatalogStackSetConstraintType, DatasourceServiceCatalogStackSetConstraint(), data, meta)
}
