// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html

package elasticbeanstalk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceElasticBeanstalkApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceElasticBeanstalkApplicationCreate,
		Read:   resourceElasticBeanstalkApplicationRead,
		Update: resourceElasticBeanstalkApplicationUpdate,
		Delete: resourceElasticBeanstalkApplicationDelete,

		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"resource_lifecycle_config": {
				Type: schema.TypeList,
				Elem: propertyApplicationApplicationResourceLifecycleConfig(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceElasticBeanstalkApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticBeanstalk::Application", data, meta)
}

func resourceElasticBeanstalkApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticBeanstalk::Application", data, meta)
}

func resourceElasticBeanstalkApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticBeanstalk::Application", data, meta)
}

func resourceElasticBeanstalkApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticBeanstalk::Application", data, meta)
}