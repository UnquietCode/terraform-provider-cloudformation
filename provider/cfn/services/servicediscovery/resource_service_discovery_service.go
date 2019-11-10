// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

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
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"health_check_custom_config": {
				Type: schema.TypeList,
				Elem: propertyServiceHealthCheckCustomConfig(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"dns_config": {
				Type: schema.TypeList,
				Elem: propertyServiceDnsConfig(),
				Required: false,
				MaxItems: 1,
			},
			"namespace_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"health_check_config": {
				Type: schema.TypeList,
				Elem: propertyServiceHealthCheckConfig(),
				Required: false,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceDiscoveryServiceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceDiscovery::Service", data, meta)
}

func resourceServiceDiscoveryServiceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceDiscovery::Service", data, meta)
}

func resourceServiceDiscoveryServiceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceDiscovery::Service", data, meta)
}

func resourceServiceDiscoveryServiceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceDiscovery::Service", data, meta)
}