// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html

package glue

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueDevEndpointType string = "AWS::Glue::DevEndpoint"

func DatasourceGlueDevEndpoint() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGlueDevEndpointRead,
		
		Schema: map[string]*schema.Schema{
			"extra_jars_s3_path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"public_key": {
				Type: schema.TypeString,
				Optional: true,
			},
			"number_of_nodes": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"arguments": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"worker_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"endpoint_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"glue_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"extra_python_libs_s3_path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_configuration": {
				Type: schema.TypeString,
				Optional: true,
			},
			"number_of_workers": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"tags": {
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

func datasourceGlueDevEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueDevEndpointType, DatasourceGlueDevEndpoint(), data, meta)
}
