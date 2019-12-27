// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html

package glue

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueJobType string = "AWS::Glue::Job"

func DatasourceGlueJob() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGlueJobRead,
		
		Schema: map[string]*schema.Schema{
			"connections": {
				Type: schema.TypeList,
				Elem: propertyJobConnectionsList(),
				Optional: true,
				MaxItems: 1,
			},
			"max_retries": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"timeout": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"allocated_capacity": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"role": {
				Type: schema.TypeString,
				Required: true,
			},
			"default_arguments": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"notification_property": {
				Type: schema.TypeList,
				Elem: propertyJobNotificationProperty(),
				Optional: true,
				MaxItems: 1,
			},
			"worker_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"log_uri": {
				Type: schema.TypeString,
				Optional: true,
			},
			"command": {
				Type: schema.TypeList,
				Elem: propertyJobJobCommand(),
				Required: true,
				MaxItems: 1,
			},
			"glue_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"execution_property": {
				Type: schema.TypeList,
				Elem: propertyJobExecutionProperty(),
				Optional: true,
				MaxItems: 1,
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
			"max_capacity": {
				Type: schema.TypeFloat,
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

func datasourceGlueJobRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueJobType, DatasourceGlueJob(), data, meta)
}
