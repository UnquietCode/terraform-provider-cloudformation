// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationCheckpointConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"configuration_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"checkpoint_interval": {
				Type: schema.TypeInt,
				Required: false,
			},
			"min_pause_between_checkpoints": {
				Type: schema.TypeInt,
				Required: false,
			},
			"checkpointing_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}