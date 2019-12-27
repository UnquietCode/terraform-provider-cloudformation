// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html

package codedeploy

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codeDeployDeploymentConfigType string = "AWS::CodeDeploy::DeploymentConfig"

func DatasourceCodeDeployDeploymentConfig() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCodeDeployDeploymentConfigRead,
		
		Schema: map[string]*schema.Schema{
			"deployment_config_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"minimum_healthy_hosts": {
				Type: schema.TypeList,
				Elem: propertyDeploymentConfigMinimumHealthyHosts(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceCodeDeployDeploymentConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codeDeployDeploymentConfigType, DatasourceCodeDeployDeploymentConfig(), data, meta)
}
