// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
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
		Exists: resourceElasticBeanstalkApplicationExists,
		Read: resourceElasticBeanstalkApplicationRead,
		Create: resourceElasticBeanstalkApplicationCreate,
		Update: resourceElasticBeanstalkApplicationUpdate,
		Delete: resourceElasticBeanstalkApplicationDelete,
		CustomizeDiff: resourceElasticBeanstalkApplicationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"resource_lifecycle_config": {
				Type: schema.TypeList,
				Elem: propertyApplicationApplicationResourceLifecycleConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceElasticBeanstalkApplicationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceElasticBeanstalkApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticBeanstalk::Application", ResourceElasticBeanstalkApplication(), data, meta)
}

func resourceElasticBeanstalkApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticBeanstalk::Application", ResourceElasticBeanstalkApplication(), data, meta)
}

func resourceElasticBeanstalkApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticBeanstalk::Application", ResourceElasticBeanstalkApplication(), data, meta)
}

func resourceElasticBeanstalkApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticBeanstalk::Application", data, meta)
}

func resourceElasticBeanstalkApplicationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::ElasticBeanstalk::Application", data, meta)
}

