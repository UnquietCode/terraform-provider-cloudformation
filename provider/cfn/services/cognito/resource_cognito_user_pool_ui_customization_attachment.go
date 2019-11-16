// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluicustomizationattachment.html

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cognitoUserPoolUICustomizationAttachmentType string = "AWS::Cognito::UserPoolUICustomizationAttachment"

var cognitoUserPoolUICustomizationAttachmentProperties map[string]string = map[string]string{
	"css": "CSS",
	"user_pool_id": "UserPoolId",
	"client_id": "ClientId",
}

func ResourceCognitoUserPoolUICustomizationAttachment() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCognitoUserPoolUICustomizationAttachmentExists,
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
			},
		},
	}
}

func resourceCognitoUserPoolUICustomizationAttachmentExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCognitoUserPoolUICustomizationAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cognitoUserPoolUICustomizationAttachmentType, ResourceCognitoUserPoolUICustomizationAttachment(), data, meta)
}

func resourceCognitoUserPoolUICustomizationAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cognitoUserPoolUICustomizationAttachmentType, ResourceCognitoUserPoolUICustomizationAttachment(), data, cognitoUserPoolUICustomizationAttachmentProperties, meta)
}

func resourceCognitoUserPoolUICustomizationAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cognitoUserPoolUICustomizationAttachmentType, ResourceCognitoUserPoolUICustomizationAttachment(), data, cognitoUserPoolUICustomizationAttachmentProperties, meta)
}

func resourceCognitoUserPoolUICustomizationAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cognitoUserPoolUICustomizationAttachmentType, data, meta)
}

func resourceCognitoUserPoolUICustomizationAttachmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cognitoUserPoolUICustomizationAttachmentType, data, meta)
}
