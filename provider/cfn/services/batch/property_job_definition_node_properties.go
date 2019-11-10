// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package batch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyJobDefinitionNodeProperties() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"main_node": {
				Type: schema.TypeInt,
				Required: true,
			},
			"node_range_properties": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionNodeRangeProperty(),
				Required: true,
			},
			"num_nodes": {
				Type: schema.TypeInt,
				Required: true,
			},
		},
	}
}