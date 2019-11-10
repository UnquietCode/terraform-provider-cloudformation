// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package codedeploy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCodeDeployDeploymentConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceCodeDeployDeploymentConfigCreate,
		Read:   resourceCodeDeployDeploymentConfigRead,
		Update: resourceCodeDeployDeploymentConfigUpdate,
		Delete: resourceCodeDeployDeploymentConfigDelete,

		Schema: map[string]*schema.Schema{
			"deployment_config_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"minimum_healthy_hosts": {
				Type: schema.TypeList,
				Elem: propertyDeploymentConfigMinimumHealthyHosts(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceCodeDeployDeploymentConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodeDeploy::DeploymentConfig", data, meta)
}

func resourceCodeDeployDeploymentConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodeDeploy::DeploymentConfig", data, meta)
}

func resourceCodeDeployDeploymentConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodeDeploy::DeploymentConfig", data, meta)
}

func resourceCodeDeployDeploymentConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodeDeploy::DeploymentConfig", data, meta)
}