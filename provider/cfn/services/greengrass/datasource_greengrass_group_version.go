// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html

package greengrass

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassGroupVersionType string = "AWS::Greengrass::GroupVersion"

func DatasourceGreengrassGroupVersion() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGreengrassGroupVersionRead,
		
		Schema: map[string]*schema.Schema{
			"logger_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"device_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"function_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"core_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"resource_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"connector_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"subscription_definition_version_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"group_id": {
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

func datasourceGreengrassGroupVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassGroupVersionType, DatasourceGreengrassGroupVersion(), data, meta)
}
