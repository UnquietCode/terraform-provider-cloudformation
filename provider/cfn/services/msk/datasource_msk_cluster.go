// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceMSKCluster() *schema.Resource {
	return &schema.Resource{
		Read: datasourceMSKClusterRead,
		
		Schema: map[string]*schema.Schema{
			"broker_node_group_info": {
				Type: schema.TypeList,
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
				Type: schema.TypeList,
				Elem: propertyClusterEncryptionInfo(),
				Optional: true,
				MaxItems: 1,
			},
			"cluster_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"client_authentication": {
				Type: schema.TypeList,
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
				Type: schema.TypeList,
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
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceMSKClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(mSKClusterType, DatasourceMSKCluster(), data, meta)
}
