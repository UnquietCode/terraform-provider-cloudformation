// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-sourcecredential.html

package codebuild

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codeBuildSourceCredentialType string = "AWS::CodeBuild::SourceCredential"

func ResourceCodeBuildSourceCredential() *schema.Resource {
	return &schema.Resource{
		Read: resourceCodeBuildSourceCredentialRead,
		Create: resourceCodeBuildSourceCredentialCreate,
		Update: resourceCodeBuildSourceCredentialUpdate,
		Delete: resourceCodeBuildSourceCredentialDelete,
		CustomizeDiff: resourceCodeBuildSourceCredentialCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"server_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"username": {
				Type: schema.TypeString,
				Optional: true,
			},
			"token": {
				Type: schema.TypeString,
				Required: true,
			},
			"auth_type": {
				Type: schema.TypeString,
				Required: true,
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

func resourceCodeBuildSourceCredentialRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codeBuildSourceCredentialType, ResourceCodeBuildSourceCredential(), data, meta)
}

func resourceCodeBuildSourceCredentialCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(codeBuildSourceCredentialType, ResourceCodeBuildSourceCredential(), data, meta)
}

func resourceCodeBuildSourceCredentialUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(codeBuildSourceCredentialType, ResourceCodeBuildSourceCredential(), data, meta)
}

func resourceCodeBuildSourceCredentialDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(codeBuildSourceCredentialType, data, meta)
}

func resourceCodeBuildSourceCredentialCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(codeBuildSourceCredentialType, data, meta)
}
