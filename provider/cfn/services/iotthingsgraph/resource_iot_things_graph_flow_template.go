// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iotthingsgraph

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoTThingsGraphFlowTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoTThingsGraphFlowTemplateCreate,
		Read:   resourceIoTThingsGraphFlowTemplateRead,
		Update: resourceIoTThingsGraphFlowTemplateUpdate,
		Delete: resourceIoTThingsGraphFlowTemplateDelete,

		Schema: map[string]*schema.Schema{
			"compatible_namespace_version": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"definition": {
				Type: schema.TypeList,
				Elem: propertyDefinitionDocument(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceIoTThingsGraphFlowTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoTThingsGraph::FlowTemplate", data, meta)
}

func resourceIoTThingsGraphFlowTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoTThingsGraph::FlowTemplate", data, meta)
}

func resourceIoTThingsGraphFlowTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoTThingsGraph::FlowTemplate", data, meta)
}

func resourceIoTThingsGraphFlowTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoTThingsGraph::FlowTemplate", data, meta)
}