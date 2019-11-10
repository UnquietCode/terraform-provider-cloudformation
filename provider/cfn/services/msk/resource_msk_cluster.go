// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html

package msk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceMSKCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceMSKClusterCreate,
		Read:   resourceMSKClusterRead,
		Update: resourceMSKClusterUpdate,
		Delete: resourceMSKClusterDelete,

		Schema: map[string]*schema.Schema{
			"broker_node_group_info": {
				Type: schema.TypeList,
				Elem: propertyClusterBrokerNodeGroupInfo(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"enhanced_monitoring": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"kafka_version": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"number_of_broker_nodes": {
				Type: schema.TypeInt,
				Required: true,
			},
			"encryption_info": {
				Type: schema.TypeList,
				Elem: propertyClusterEncryptionInfo(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"cluster_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"client_authentication": {
				Type: schema.TypeList,
				Elem: propertyClusterClientAuthentication(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
			"configuration_info": {
				Type: schema.TypeList,
				Elem: propertyClusterConfigurationInfo(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceMSKClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::MSK::Cluster", data, meta)
}

func resourceMSKClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::MSK::Cluster", data, meta)
}

func resourceMSKClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::MSK::Cluster", data, meta)
}

func resourceMSKClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::MSK::Cluster", data, meta)
}