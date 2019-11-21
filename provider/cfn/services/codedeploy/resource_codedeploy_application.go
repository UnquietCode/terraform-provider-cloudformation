// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-application.html

package codedeploy

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codeDeployApplicationType string = "AWS::CodeDeploy::Application"

func ResourceCodeDeployApplication() *schema.Resource {
	return &schema.Resource{
		Read: resourceCodeDeployApplicationRead,
		Create: resourceCodeDeployApplicationCreate,
		Update: resourceCodeDeployApplicationUpdate,
		Delete: resourceCodeDeployApplicationDelete,
		CustomizeDiff: resourceCodeDeployApplicationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"compute_platform": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceCodeDeployApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codeDeployApplicationType, ResourceCodeDeployApplication(), data, meta)
}

func resourceCodeDeployApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(codeDeployApplicationType, ResourceCodeDeployApplication(), data, meta)
}

func resourceCodeDeployApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(codeDeployApplicationType, ResourceCodeDeployApplication(), data, meta)
}

func resourceCodeDeployApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(codeDeployApplicationType, data, meta)
}

func resourceCodeDeployApplicationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(codeDeployApplicationType, data, meta)
}
