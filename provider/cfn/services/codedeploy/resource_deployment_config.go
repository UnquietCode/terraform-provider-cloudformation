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

func ResourceDeploymentConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeploymentConfigCreate,
		Read:   resourceDeploymentConfigRead,
		Update: resourceDeploymentConfigUpdate,
		Delete: resourceDeploymentConfigDelete,

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

func resourceDeploymentConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodeDeploy::DeploymentConfig", data, meta)
}

func resourceDeploymentConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodeDeploy::DeploymentConfig", data, meta)
}

func resourceDeploymentConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodeDeploy::DeploymentConfig", data, meta)
}

func resourceDeploymentConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodeDeploy::DeploymentConfig", data, meta)
}