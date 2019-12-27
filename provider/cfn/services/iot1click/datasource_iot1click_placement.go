// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html

package iot1click

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioT1ClickPlacementType string = "AWS::IoT1Click::Placement"

func DatasourceIoT1ClickPlacement() *schema.Resource {
	return &schema.Resource{
		Read: datasourceIoT1ClickPlacementRead,
		
		Schema: map[string]*schema.Schema{
			"placement_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"project_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"associated_devices": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"attributes": {
				Type: schema.TypeString,
				Optional: true,
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

func datasourceIoT1ClickPlacementRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioT1ClickPlacementType, DatasourceIoT1ClickPlacement(), data, meta)
}
