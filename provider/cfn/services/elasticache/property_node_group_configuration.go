// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package elasticache

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyNodeGroupConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_group_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"primary_availability_zone": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"replica_availability_zones": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"replica_count": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"slots": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}