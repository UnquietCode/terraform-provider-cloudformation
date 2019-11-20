// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html

package codedeploy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codeDeployDeploymentConfigType string = "AWS::CodeDeploy::DeploymentConfig"

func ResourceCodeDeployDeploymentConfig() *schema.Resource {
	return &schema.Resource{
		Read: resourceCodeDeployDeploymentConfigRead,
		Create: resourceCodeDeployDeploymentConfigCreate,
		Update: resourceCodeDeployDeploymentConfigUpdate,
		Delete: resourceCodeDeployDeploymentConfigDelete,
		CustomizeDiff: resourceCodeDeployDeploymentConfigCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"deployment_config_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"minimum_healthy_hosts": {
				Type: schema.TypeSet,
				Elem: propertyDeploymentConfigMinimumHealthyHosts(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCodeDeployDeploymentConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codeDeployDeploymentConfigType, ResourceCodeDeployDeploymentConfig(), data, meta)
}

func resourceCodeDeployDeploymentConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(codeDeployDeploymentConfigType, ResourceCodeDeployDeploymentConfig(), data, meta)
}

func resourceCodeDeployDeploymentConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(codeDeployDeploymentConfigType, ResourceCodeDeployDeploymentConfig(), data, meta)
}

func resourceCodeDeployDeploymentConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(codeDeployDeploymentConfigType, data, meta)
}

func resourceCodeDeployDeploymentConfigCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(codeDeployDeploymentConfigType, data, meta)
}
