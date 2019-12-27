// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-servicelinkedrole.html

package iam

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const iAMServiceLinkedRoleType string = "AWS::IAM::ServiceLinkedRole"

func DatasourceIAMServiceLinkedRole() *schema.Resource {
	return &schema.Resource{
		Read: datasourceIAMServiceLinkedRoleRead,
		
		Schema: map[string]*schema.Schema{
			"custom_suffix": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"aws_service_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceIAMServiceLinkedRoleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(iAMServiceLinkedRoleType, DatasourceIAMServiceLinkedRole(), data, meta)
}
