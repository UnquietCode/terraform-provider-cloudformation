// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html

package elasticbeanstalk

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticBeanstalkApplicationType string = "AWS::ElasticBeanstalk::Application"

func DatasourceElasticBeanstalkApplication() *schema.Resource {
	return &schema.Resource{
		Read: datasourceElasticBeanstalkApplicationRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceElasticBeanstalkApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticBeanstalkApplicationType, DatasourceElasticBeanstalkApplication(), data, meta)
}
