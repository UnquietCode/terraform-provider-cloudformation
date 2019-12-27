// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html

package iot1click

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioT1ClickProjectType string = "AWS::IoT1Click::Project"

func DatasourceIoT1ClickProject() *schema.Resource {
	return &schema.Resource{
		Read: datasourceIoT1ClickProjectRead,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"placement_template": {
				Type: schema.TypeList,
				Elem: propertyProjectPlacementTemplate(),
				Required: true,
				MaxItems: 1,
			},
			"project_name": {
				Type: schema.TypeString,
				Optional: true,
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

func datasourceIoT1ClickProjectRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioT1ClickProjectType, DatasourceIoT1ClickProject(), data, meta)
}
