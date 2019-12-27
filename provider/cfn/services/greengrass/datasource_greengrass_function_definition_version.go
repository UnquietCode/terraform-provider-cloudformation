// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html

package greengrass

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassFunctionDefinitionVersionType string = "AWS::Greengrass::FunctionDefinitionVersion"

func DatasourceGreengrassFunctionDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGreengrassFunctionDefinitionVersionRead,
		
		Schema: map[string]*schema.Schema{
			"default_config": {
				Type: schema.TypeList,
				Elem: propertyFunctionDefinitionVersionDefaultConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"functions": {
				Type: schema.TypeList,
				Elem: propertyFunctionDefinitionVersionFunction(),
				Required: true,
			},
			"function_definition_id": {
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

func datasourceGreengrassFunctionDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassFunctionDefinitionVersionType, DatasourceGreengrassFunctionDefinitionVersion(), data, meta)
}
