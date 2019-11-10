// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package codepipeline

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyPipelineStageTransition() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"reason": {
				Type: schema.TypeString,
				Required: true,
			},
			"stage_name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}