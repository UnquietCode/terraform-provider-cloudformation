// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotthingsgraph-flowtemplate.html

package iotthingsgraph

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTThingsGraphFlowTemplateType string = "AWS::IoTThingsGraph::FlowTemplate"

func DatasourceIoTThingsGraphFlowTemplate() *schema.Resource {
	return &schema.Resource{
		Read: datasourceIoTThingsGraphFlowTemplateRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceIoTThingsGraphFlowTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTThingsGraphFlowTemplateType, DatasourceIoTThingsGraphFlowTemplate(), data, meta)
}
