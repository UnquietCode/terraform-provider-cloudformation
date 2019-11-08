// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyContainerDefinition() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"container_hostname": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"environment": {
				Type: schema.TypeMap,
				Required: false,
				ForceNew: true,
			},
			"model_data_url": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"image": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}