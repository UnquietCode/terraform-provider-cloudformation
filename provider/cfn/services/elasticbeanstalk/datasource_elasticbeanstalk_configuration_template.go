// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html

package elasticbeanstalk

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticBeanstalkConfigurationTemplateType string = "AWS::ElasticBeanstalk::ConfigurationTemplate"

func DatasourceElasticBeanstalkConfigurationTemplate() *schema.Resource {
	return &schema.Resource{
		Read: datasourceElasticBeanstalkConfigurationTemplateRead,
		
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"environment_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"option_settings": {
				Type: schema.TypeList,
				Elem: propertyConfigurationTemplateConfigurationOptionSetting(),
				Optional: true,
			},
			"platform_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"solution_stack_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_configuration": {
				Type: schema.TypeList,
				Elem: propertyConfigurationTemplateSourceConfiguration(),
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

func datasourceElasticBeanstalkConfigurationTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticBeanstalkConfigurationTemplateType, DatasourceElasticBeanstalkConfigurationTemplate(), data, meta)
}
