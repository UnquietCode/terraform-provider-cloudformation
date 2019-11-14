// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-sourcecredential.html

package codebuild

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCodeBuildSourceCredential() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCodeBuildSourceCredentialExists,
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
			},
		},
	}
}

func resourceCodeBuildSourceCredentialExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCodeBuildSourceCredentialRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodeBuild::SourceCredential", ResourceCodeBuildSourceCredential(), data, meta)
}

func resourceCodeBuildSourceCredentialCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodeBuild::SourceCredential", ResourceCodeBuildSourceCredential(), data, meta)
}

func resourceCodeBuildSourceCredentialUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodeBuild::SourceCredential", ResourceCodeBuildSourceCredential(), data, meta)
}

func resourceCodeBuildSourceCredentialDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodeBuild::SourceCredential", data, meta)
}

func resourceCodeBuildSourceCredentialCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
