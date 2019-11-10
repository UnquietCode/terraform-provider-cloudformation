// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyResourceDefinitionVersionSageMakerMachineLearningModelResourceData() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"destination_path": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sage_maker_job_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}