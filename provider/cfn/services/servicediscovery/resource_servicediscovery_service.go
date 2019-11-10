// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html

package servicediscovery

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceDiscoveryService() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDiscoveryServiceCreate,
		Read:   resourceServiceDiscoveryServiceRead,
		Update: resourceServiceDiscoveryServiceUpdate,
		Delete: resourceServiceDiscoveryServiceDelete,

		Schema: map[string]*schema.Schema{
			"the_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"health_check_custom_config": {
				Type: schema.TypeList,
				Elem: propertyServiceHealthCheckCustomConfig(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"dns_config": {
				Type: schema.TypeList,
				Elem: propertyServiceDnsConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"namespace_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"health_check_config": {
				Type: schema.TypeList,
				Elem: propertyServiceHealthCheckConfig(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceServiceDiscoveryServiceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceDiscovery::Service", ResourceServiceDiscoveryService(), data, meta)
}

func resourceServiceDiscoveryServiceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceDiscovery::Service", ResourceServiceDiscoveryService(), data, meta)
}

func resourceServiceDiscoveryServiceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceDiscovery::Service", ResourceServiceDiscoveryService(), data, meta)
}

func resourceServiceDiscoveryServiceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceDiscovery::Service", data, meta)
}