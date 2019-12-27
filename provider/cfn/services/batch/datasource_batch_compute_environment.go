// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html

package batch

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const batchComputeEnvironmentType string = "AWS::Batch::ComputeEnvironment"

func DatasourceBatchComputeEnvironment() *schema.Resource {
	return &schema.Resource{
		Read: datasourceBatchComputeEnvironmentRead,
		
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"service_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"compute_environment_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"compute_resources": {
				Type: schema.TypeList,
				Elem: propertyComputeEnvironmentComputeResources(),
				Optional: true,
				MaxItems: 1,
			},
			"state": {
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

func datasourceBatchComputeEnvironmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(batchComputeEnvironmentType, DatasourceBatchComputeEnvironment(), data, meta)
}
