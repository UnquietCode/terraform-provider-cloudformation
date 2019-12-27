// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html

package batch

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const batchJobDefinitionType string = "AWS::Batch::JobDefinition"

func DatasourceBatchJobDefinition() *schema.Resource {
	return &schema.Resource{
		Read: datasourceBatchJobDefinitionRead,
		
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"parameters": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"node_properties": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionNodeProperties(),
				Optional: true,
				MaxItems: 1,
			},
			"timeout": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionTimeout(),
				Optional: true,
				MaxItems: 1,
			},
			"container_properties": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionContainerProperties(),
				Optional: true,
				MaxItems: 1,
			},
			"job_definition_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"retry_strategy": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionRetryStrategy(),
				Optional: true,
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

func datasourceBatchJobDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(batchJobDefinitionType, DatasourceBatchJobDefinition(), data, meta)
}
