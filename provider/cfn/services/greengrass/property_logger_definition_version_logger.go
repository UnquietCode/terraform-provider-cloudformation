// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyLoggerDefinitionVersionLogger() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"space": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"level": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"component": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}