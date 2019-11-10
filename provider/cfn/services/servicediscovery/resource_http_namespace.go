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

func ResourceHttpNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceHttpNamespaceCreate,
		Read:   resourceHttpNamespaceRead,
		Update: resourceHttpNamespaceUpdate,
		Delete: resourceHttpNamespaceDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceHttpNamespaceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceDiscovery::HttpNamespace", data, meta)
}

func resourceHttpNamespaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceDiscovery::HttpNamespace", data, meta)
}

func resourceHttpNamespaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceDiscovery::HttpNamespace", data, meta)
}

func resourceHttpNamespaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceDiscovery::HttpNamespace", data, meta)
}