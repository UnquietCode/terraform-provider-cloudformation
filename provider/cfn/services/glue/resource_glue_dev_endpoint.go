// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueDevEndpointType string = "AWS::Glue::DevEndpoint"

var glueDevEndpointProperties map[string]string = map[string]string{
	"extra_jars_s3_path": "ExtraJarsS3Path",
	"public_key": "PublicKey",
	"number_of_nodes": "NumberOfNodes",
	"arguments": "Arguments",
	"subnet_id": "SubnetId",
	"security_group_ids": "SecurityGroupIds",
	"role_arn": "RoleArn",
	"worker_type": "WorkerType",
	"endpoint_name": "EndpointName",
	"glue_version": "GlueVersion",
	"extra_python_libs_s3_path": "ExtraPythonLibsS3Path",
	"security_configuration": "SecurityConfiguration",
	"number_of_workers": "NumberOfWorkers",
	"tags": "Tags",
}

func ResourceGlueDevEndpoint() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGlueDevEndpointExists,
		Read: resourceGlueDevEndpointRead,
		Create: resourceGlueDevEndpointCreate,
		Update: resourceGlueDevEndpointUpdate,
		Delete: resourceGlueDevEndpointDelete,
		CustomizeDiff: resourceGlueDevEndpointCustomizeDiff,
		
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
				Type: schema.TypeMap,
				Optional: true,
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
				Type: schema.TypeMap,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGlueDevEndpointExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGlueDevEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueDevEndpointType, ResourceGlueDevEndpoint(), data, meta)
}

func resourceGlueDevEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(glueDevEndpointType, ResourceGlueDevEndpoint(), data, glueDevEndpointProperties, meta)
}

func resourceGlueDevEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(glueDevEndpointType, ResourceGlueDevEndpoint(), data, glueDevEndpointProperties, meta)
}

func resourceGlueDevEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(glueDevEndpointType, data, meta)
}

func resourceGlueDevEndpointCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(glueDevEndpointType, data, meta)
}
