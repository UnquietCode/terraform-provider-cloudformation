// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html

package ecs

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var serviceDeploymentConfigurationProperties map[string]string = map[string]string{
	"maximum_percent": "MaximumPercent",
	"minimum_healthy_percent": "MinimumHealthyPercent",
}

func propertyServiceDeploymentConfiguration(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"maximum_percent": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"minimum_healthy_percent": {
				Type: schema.TypeInt,
				Optional: true,
			},
		},
	}
}
