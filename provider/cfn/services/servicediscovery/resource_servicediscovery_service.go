// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html

package servicediscovery

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceDiscoveryServiceType string = "AWS::ServiceDiscovery::Service"

var serviceDiscoveryServiceProperties map[string]string = map[string]string{
	"description": "Description",
	"health_check_custom_config": "HealthCheckCustomConfig",
	"dns_config": "DnsConfig",
	"namespace_id": "NamespaceId",
	"health_check_config": "HealthCheckConfig",
	"name": "Name",
}

func ResourceServiceDiscoveryService() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceDiscoveryServiceExists,
		Read: resourceServiceDiscoveryServiceRead,
		Create: resourceServiceDiscoveryServiceCreate,
		Update: resourceServiceDiscoveryServiceUpdate,
		Delete: resourceServiceDiscoveryServiceDelete,
		CustomizeDiff: resourceServiceDiscoveryServiceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"health_check_custom_config": {
				Type: schema.TypeList,
				Elem: propertyServiceHealthCheckCustomConfig(),
				Optional: true,
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
			},
			"health_check_config": {
				Type: schema.TypeList,
				Elem: propertyServiceHealthCheckConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceDiscoveryServiceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceDiscoveryServiceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceDiscoveryServiceType, ResourceServiceDiscoveryService(), data, meta)
}

func resourceServiceDiscoveryServiceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceDiscoveryServiceType, ResourceServiceDiscoveryService(), data, serviceDiscoveryServiceProperties, meta)
}

func resourceServiceDiscoveryServiceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceDiscoveryServiceType, ResourceServiceDiscoveryService(), data, serviceDiscoveryServiceProperties, meta)
}

func resourceServiceDiscoveryServiceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceDiscoveryServiceType, data, meta)
}

func resourceServiceDiscoveryServiceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceDiscoveryServiceType, data, meta)
}
