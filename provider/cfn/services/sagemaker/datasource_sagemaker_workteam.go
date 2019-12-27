// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workteam.html

package sagemaker

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sageMakerWorkteamType string = "AWS::SageMaker::Workteam"

func DatasourceSageMakerWorkteam() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSageMakerWorkteamRead,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"notification_configuration": {
				Type: schema.TypeList,
				Elem: propertyWorkteamNotificationConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"workteam_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"member_definitions": {
				Type: schema.TypeList,
				Elem: propertyWorkteamMemberDefinition(),
				Optional: true,
			},
			"tags": misc.PropertyTags(),
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

func datasourceSageMakerWorkteamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sageMakerWorkteamType, DatasourceSageMakerWorkteam(), data, meta)
}
