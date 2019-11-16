// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html

package elasticbeanstalk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticBeanstalkApplicationVersionType string = "AWS::ElasticBeanstalk::ApplicationVersion"

var elasticBeanstalkApplicationVersionProperties map[string]string = map[string]string{
	"application_name": "ApplicationName",
	"description": "Description",
	"source_bundle": "SourceBundle",
}

func ResourceElasticBeanstalkApplicationVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceElasticBeanstalkApplicationVersionExists,
		Read: resourceElasticBeanstalkApplicationVersionRead,
		Create: resourceElasticBeanstalkApplicationVersionCreate,
		Update: resourceElasticBeanstalkApplicationVersionUpdate,
		Delete: resourceElasticBeanstalkApplicationVersionDelete,
		CustomizeDiff: resourceElasticBeanstalkApplicationVersionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_bundle": {
				Type: schema.TypeSet,
				Elem: propertyApplicationVersionSourceBundle(),
				Required: true,
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

func resourceElasticBeanstalkApplicationVersionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceElasticBeanstalkApplicationVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticBeanstalkApplicationVersionType, ResourceElasticBeanstalkApplicationVersion(), data, meta)
}

func resourceElasticBeanstalkApplicationVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(elasticBeanstalkApplicationVersionType, ResourceElasticBeanstalkApplicationVersion(), data, elasticBeanstalkApplicationVersionProperties, meta)
}

func resourceElasticBeanstalkApplicationVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(elasticBeanstalkApplicationVersionType, ResourceElasticBeanstalkApplicationVersion(), data, elasticBeanstalkApplicationVersionProperties, meta)
}

func resourceElasticBeanstalkApplicationVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(elasticBeanstalkApplicationVersionType, data, meta)
}

func resourceElasticBeanstalkApplicationVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(elasticBeanstalkApplicationVersionType, data, meta)
}
