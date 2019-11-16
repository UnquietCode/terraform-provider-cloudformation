// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotthingsgraph-flowtemplate.html

package iotthingsgraph

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTThingsGraphFlowTemplateType string = "AWS::IoTThingsGraph::FlowTemplate"

var ioTThingsGraphFlowTemplateProperties map[string]string = map[string]string{
	"compatible_namespace_version": "CompatibleNamespaceVersion",
	"definition": "Definition",
}

func ResourceIoTThingsGraphFlowTemplate() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIoTThingsGraphFlowTemplateExists,
		Read: resourceIoTThingsGraphFlowTemplateRead,
		Create: resourceIoTThingsGraphFlowTemplateCreate,
		Update: resourceIoTThingsGraphFlowTemplateUpdate,
		Delete: resourceIoTThingsGraphFlowTemplateDelete,
		CustomizeDiff: resourceIoTThingsGraphFlowTemplateCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"compatible_namespace_version": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"definition": {
				Type: schema.TypeList,
				Elem: propertyFlowTemplateDefinitionDocument(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIoTThingsGraphFlowTemplateExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIoTThingsGraphFlowTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTThingsGraphFlowTemplateType, ResourceIoTThingsGraphFlowTemplate(), data, meta)
}

func resourceIoTThingsGraphFlowTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTThingsGraphFlowTemplateType, ResourceIoTThingsGraphFlowTemplate(), data, ioTThingsGraphFlowTemplateProperties, meta)
}

func resourceIoTThingsGraphFlowTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTThingsGraphFlowTemplateType, ResourceIoTThingsGraphFlowTemplate(), data, ioTThingsGraphFlowTemplateProperties, meta)
}

func resourceIoTThingsGraphFlowTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTThingsGraphFlowTemplateType, data, meta)
}

func resourceIoTThingsGraphFlowTemplateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTThingsGraphFlowTemplateType, data, meta)
}
