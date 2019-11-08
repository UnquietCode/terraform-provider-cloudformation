// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package servicediscovery

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceDiscoveryInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDiscoveryInstanceCreate,
		Read:   resourceServiceDiscoveryInstanceRead,
		Update: resourceServiceDiscoveryInstanceUpdate,
		Delete: resourceServiceDiscoveryInstanceDelete,

		Schema: map[string]*schema.Schema{
			"instance_attributes": {
				Type: schema.TypeMap,
				Required: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"service_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceDiscoveryInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceDiscovery::Instance", data, meta)
}

func resourceServiceDiscoveryInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceDiscovery::Instance", data, meta)
}

func resourceServiceDiscoveryInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceDiscovery::Instance", data, meta)
}

func resourceServiceDiscoveryInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceDiscovery::Instance", data, meta)
}