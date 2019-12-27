// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-resourcepolicy.html

package secretsmanager

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const secretsManagerResourcePolicyType string = "AWS::SecretsManager::ResourcePolicy"

func DatasourceSecretsManagerResourcePolicy() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSecretsManagerResourcePolicyRead,
		
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"resource_policy": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.ValidateJsonString,
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

func datasourceSecretsManagerResourcePolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(secretsManagerResourcePolicyType, DatasourceSecretsManagerResourcePolicy(), data, meta)
}
