// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html

package msk

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const mSKClusterType string = "AWS::MSK::Cluster"

func ResourceMSKCluster() *schema.Resource {
	return &schema.Resource{
		Read: resourceMSKClusterRead,
		Create: resourceMSKClusterCreate,
		Update: resourceMSKClusterUpdate,
		Delete: resourceMSKClusterDelete,
		CustomizeDiff: resourceMSKClusterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"broker_node_group_info": {
				Type: schema.TypeSet,
				Elem: propertyClusterBrokerNodeGroupInfo(),
				Required: true,
				MaxItems: 1,
			},
			"enhanced_monitoring": {
				Type: schema.TypeString,
				Optional: true,
			},
			"kafka_version": {
				Type: schema.TypeString,
				Required: true,
			},
			"number_of_broker_nodes": {
				Type: schema.TypeInt,
				Required: true,
			},
			"encryption_info": {
				Type: schema.TypeSet,
				Elem: propertyClusterEncryptionInfo(),
				Optional: true,
				MaxItems: 1,
			},
			"cluster_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"client_authentication": {
				Type: schema.TypeSet,
				Elem: propertyClusterClientAuthentication(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"configuration_info": {
				Type: schema.TypeSet,
				Elem: propertyClusterConfigurationInfo(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceMSKClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(mSKClusterType, ResourceMSKCluster(), data, meta)
}

func resourceMSKClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(mSKClusterType, ResourceMSKCluster(), data, meta)
}

func resourceMSKClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(mSKClusterType, ResourceMSKCluster(), data, meta)
}

func resourceMSKClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(mSKClusterType, data, meta)
}

func resourceMSKClusterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(mSKClusterType, data, meta)
}
