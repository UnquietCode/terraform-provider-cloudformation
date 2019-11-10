// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-httpnamespace.html

package servicediscovery

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceDiscoveryHttpNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDiscoveryHttpNamespaceCreate,
		Read:   resourceServiceDiscoveryHttpNamespaceRead,
		Delete: resourceServiceDiscoveryHttpNamespaceDelete,

		Schema: map[string]*schema.Schema{
			"the_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceServiceDiscoveryHttpNamespaceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceDiscovery::HttpNamespace", ResourceServiceDiscoveryHttpNamespace(), data, meta)
}

func resourceServiceDiscoveryHttpNamespaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceDiscovery::HttpNamespace", ResourceServiceDiscoveryHttpNamespace(), data, meta)
}

func resourceServiceDiscoveryHttpNamespaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceDiscovery::HttpNamespace", ResourceServiceDiscoveryHttpNamespace(), data, meta)
}

func resourceServiceDiscoveryHttpNamespaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceDiscovery::HttpNamespace", data, meta)
}