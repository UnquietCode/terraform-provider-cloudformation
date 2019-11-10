// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyClusterStepConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action_on_failure": {
				Type: schema.TypeString,
				Required: false,
			},
			"hadoop_jar_step": {
				Type: schema.TypeList,
				Elem: propertyClusterHadoopJarStepConfig(),
				Required: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}