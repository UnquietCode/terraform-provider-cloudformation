// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html

package cognito

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var identityPoolRoleAttachmentRoleMappingProperties map[string]string = map[string]string{
	"type": "Type",
	"ambiguous_role_resolution": "AmbiguousRoleResolution",
	"rules_configuration": "RulesConfiguration",
	"identity_provider": "IdentityProvider",
}

func propertyIdentityPoolRoleAttachmentRoleMapping(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"ambiguous_role_resolution": {
				Type: schema.TypeString,
				Optional: true,
			},
			"rules_configuration": {
				Type: schema.TypeList,
				Elem: propertyIdentityPoolRoleAttachmentRulesConfigurationType(),
				Optional: true,
				MaxItems: 1,
			},
			"identity_provider": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
