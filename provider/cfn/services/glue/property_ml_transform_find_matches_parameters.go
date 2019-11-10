// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyMLTransformFindMatchesParameters() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"precision_recall_tradeoff": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"enforce_provided_labels": {
				Type: schema.TypeBool,
				Required: false,
			},
			"primary_key_column_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"accuracy_cost_tradeoff": {
				Type: schema.TypeFloat,
				Required: false,
			},
		},
	}
}