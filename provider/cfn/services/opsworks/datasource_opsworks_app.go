// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html

package opsworks

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const opsWorksAppType string = "AWS::OpsWorks::App"

func DatasourceOpsWorksApp() *schema.Resource {
	return &schema.Resource{
		Read: datasourceOpsWorksAppRead,
		
		Schema: map[string]*schema.Schema{
			"app_source": {
				Type: schema.TypeList,
				Elem: propertyAppSource(),
				Optional: true,
				MaxItems: 1,
			},
			"attributes": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"data_sources": {
				Type: schema.TypeSet,
				Elem: propertyAppDataSource(),
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"domains": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"enable_ssl": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"environment": {
				Type: schema.TypeList,
				Elem: propertyAppEnvironmentVariable(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"shortname": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ssl_configuration": {
				Type: schema.TypeList,
				Elem: propertyAppSslConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"stack_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
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

func datasourceOpsWorksAppRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(opsWorksAppType, DatasourceOpsWorksApp(), data, meta)
}
