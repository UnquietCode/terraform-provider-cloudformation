// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluicustomizationattachment.html

package cognito

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cognitoUserPoolUICustomizationAttachmentType string = "AWS::Cognito::UserPoolUICustomizationAttachment"

func ResourceCognitoUserPoolUICustomizationAttachment() *schema.Resource {
	return &schema.Resource{
		Read: resourceCognitoUserPoolUICustomizationAttachmentRead,
		Create: resourceCognitoUserPoolUICustomizationAttachmentCreate,
		Update: resourceCognitoUserPoolUICustomizationAttachmentUpdate,
		Delete: resourceCognitoUserPoolUICustomizationAttachmentDelete,
		CustomizeDiff: resourceCognitoUserPoolUICustomizationAttachmentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"css": {
				Type: schema.TypeString,
				Optional: true,
			},
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"client_id": {
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

func resourceCognitoUserPoolUICustomizationAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cognitoUserPoolUICustomizationAttachmentType, ResourceCognitoUserPoolUICustomizationAttachment(), data, meta)
}

func resourceCognitoUserPoolUICustomizationAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cognitoUserPoolUICustomizationAttachmentType, ResourceCognitoUserPoolUICustomizationAttachment(), data, meta)
}

func resourceCognitoUserPoolUICustomizationAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cognitoUserPoolUICustomizationAttachmentType, ResourceCognitoUserPoolUICustomizationAttachment(), data, meta)
}

func resourceCognitoUserPoolUICustomizationAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cognitoUserPoolUICustomizationAttachmentType, data, meta)
}

func resourceCognitoUserPoolUICustomizationAttachmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cognitoUserPoolUICustomizationAttachmentType, data, meta)
}
