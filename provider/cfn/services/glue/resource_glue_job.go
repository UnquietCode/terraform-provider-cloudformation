// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGlueJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlueJobCreate,
		Read:   resourceGlueJobRead,
		Update: resourceGlueJobUpdate,
		Delete: resourceGlueJobDelete,

		Schema: map[string]*schema.Schema{
			"connections": {
				Type: schema.TypeList,
				Elem: propertyConnectionsList(),
				Required: false,
				MaxItems: 1,
			},
			"max_retries": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"timeout": {
				Type: schema.TypeInt,
				Required: false,
			},
			"allocated_capacity": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"role": {
				Type: schema.TypeString,
				Required: true,
			},
			"default_arguments": {
				Type: schema.TypeMap,
				Required: false,
			},
			"notification_property": {
				Type: schema.TypeList,
				Elem: propertyNotificationProperty(),
				Required: false,
				MaxItems: 1,
			},
			"worker_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"log_uri": {
				Type: schema.TypeString,
				Required: false,
			},
			"command": {
				Type: schema.TypeList,
				Elem: propertyJobCommand(),
				Required: true,
				MaxItems: 1,
			},
			"glue_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"execution_property": {
				Type: schema.TypeList,
				Elem: propertyExecutionProperty(),
				Required: false,
				MaxItems: 1,
			},
			"security_configuration": {
				Type: schema.TypeString,
				Required: false,
			},
			"number_of_workers": {
				Type: schema.TypeInt,
				Required: false,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"max_capacity": {
				Type: schema.TypeFloat,
				Required: false,
			},
		},
	}
}

func resourceGlueJobCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::Job", data, meta)
}

func resourceGlueJobRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::Job", data, meta)
}

func resourceGlueJobUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::Job", data, meta)
}

func resourceGlueJobDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::Job", data, meta)
}