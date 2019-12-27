// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html

package iotevents

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTEventsInputType string = "AWS::IoTEvents::Input"

func DatasourceIoTEventsInput() *schema.Resource {
	return &schema.Resource{
		Read: datasourceIoTEventsInputRead,
		
		Schema: map[string]*schema.Schema{
			"input_definition": {
				Type: schema.TypeList,
				Elem: propertyInputInputDefinition(),
				Optional: true,
				MaxItems: 1,
			},
			"input_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"input_description": {
				Type: schema.TypeString,
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

func datasourceIoTEventsInputRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTEventsInputType, DatasourceIoTEventsInput(), data, meta)
}
