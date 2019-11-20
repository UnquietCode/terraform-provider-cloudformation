// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html

package elasticbeanstalk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticBeanstalkApplicationType string = "AWS::ElasticBeanstalk::Application"

func ResourceElasticBeanstalkApplication() *schema.Resource {
	return &schema.Resource{
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
				Type: schema.TypeSet,
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

func resourceElasticBeanstalkApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticBeanstalkApplicationType, ResourceElasticBeanstalkApplication(), data, meta)
}

func resourceElasticBeanstalkApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(elasticBeanstalkApplicationType, ResourceElasticBeanstalkApplication(), data, meta)
}

func resourceElasticBeanstalkApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(elasticBeanstalkApplicationType, ResourceElasticBeanstalkApplication(), data, meta)
}

func resourceElasticBeanstalkApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(elasticBeanstalkApplicationType, data, meta)
}

func resourceElasticBeanstalkApplicationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(elasticBeanstalkApplicationType, data, meta)
}
