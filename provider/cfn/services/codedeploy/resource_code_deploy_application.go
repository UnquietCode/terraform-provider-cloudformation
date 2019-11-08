// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package codedeploy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCodeDeployApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceCodeDeployApplicationCreate,
		Read:   resourceCodeDeployApplicationRead,
		Update: resourceCodeDeployApplicationUpdate,
		Delete: resourceCodeDeployApplicationDelete,

		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"compute_platform": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceCodeDeployApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodeDeploy::Application", data, meta)
}

func resourceCodeDeployApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodeDeploy::Application", data, meta)
}

func resourceCodeDeployApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodeDeploy::Application", data, meta)
}

func resourceCodeDeployApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodeDeploy::Application", data, meta)
}