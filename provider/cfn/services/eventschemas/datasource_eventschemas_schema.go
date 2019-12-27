// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html

package eventschemas

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eventSchemasSchemaType string = "AWS::EventSchemas::Schema"

func DatasourceEventSchemasSchema() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEventSchemasSchemaRead,
		
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"content": {
				Type: schema.TypeString,
				Required: true,
			},
			"registry_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"schema_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertySchemaTagsEntry(),
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

func datasourceEventSchemasSchemaRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eventSchemasSchemaType, DatasourceEventSchemasSchema(), data, meta)
}
