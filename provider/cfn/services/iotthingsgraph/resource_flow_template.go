// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iotthingsgraph

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceFlowTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceFlowTemplateCreate,
		Read:   resourceFlowTemplateRead,
		Update: resourceFlowTemplateUpdate,
		Delete: resourceFlowTemplateDelete,

		Schema: map[string]*schema.Schema{
			"compatible_namespace_version": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"definition": {
				Type: schema.TypeList,
				Elem: propertyFlowTemplateDefinitionDocument(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceFlowTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoTThingsGraph::FlowTemplate", data, meta)
}

func resourceFlowTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoTThingsGraph::FlowTemplate", data, meta)
}

func resourceFlowTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoTThingsGraph::FlowTemplate", data, meta)
}

func resourceFlowTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoTThingsGraph::FlowTemplate", data, meta)
}