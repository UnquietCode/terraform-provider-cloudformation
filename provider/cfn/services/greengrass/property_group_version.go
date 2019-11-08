// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyGroupVersion() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"logger_definition_version_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"device_definition_version_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"function_definition_version_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"core_definition_version_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"resource_definition_version_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"connector_definition_version_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"subscription_definition_version_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}