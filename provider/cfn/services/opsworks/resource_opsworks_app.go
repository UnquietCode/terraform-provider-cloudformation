// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html

package opsworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceOpsWorksApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceOpsWorksAppCreate,
		Read:   resourceOpsWorksAppRead,
		Update: resourceOpsWorksAppUpdate,
		Delete: resourceOpsWorksAppDelete,

		Schema: map[string]*schema.Schema{
			"app_source": {
				Type: schema.TypeList,
				Elem: propertyAppSource(),
				Required: false,
				MaxItems: 1,
			},
			"attributes": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"data_sources": {
				Type: schema.TypeSet,
				Elem: propertyAppDataSource(),
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"domains": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"enable_ssl": {
				Type: schema.TypeBool,
				Required: false,
			},
			"environment": {
				Type: schema.TypeList,
				Elem: propertyAppEnvironmentVariable(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"shortname": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"ssl_configuration": {
				Type: schema.TypeList,
				Elem: propertyAppSslConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"stack_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceOpsWorksAppCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::OpsWorks::App", data, meta)
}

func resourceOpsWorksAppRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::OpsWorks::App", data, meta)
}

func resourceOpsWorksAppUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::OpsWorks::App", data, meta)
}

func resourceOpsWorksAppDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::OpsWorks::App", data, meta)
}