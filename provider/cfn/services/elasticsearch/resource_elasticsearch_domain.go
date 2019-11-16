// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html

package elasticsearch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticsearchDomainType string = "AWS::Elasticsearch::Domain"

func ResourceElasticsearchDomain() *schema.Resource {
	return &schema.Resource{
		Exists: resourceElasticsearchDomainExists,
		Read: resourceElasticsearchDomainRead,
		Create: resourceElasticsearchDomainCreate,
		Update: resourceElasticsearchDomainUpdate,
		Delete: resourceElasticsearchDomainDelete,
		CustomizeDiff: resourceElasticsearchDomainCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"access_policies": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"advanced_options": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"domain_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ebs_options": {
				Type: schema.TypeSet,
				Elem: propertyDomainEBSOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"elasticsearch_cluster_config": {
				Type: schema.TypeSet,
				Elem: propertyDomainElasticsearchClusterConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"elasticsearch_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"encryption_at_rest_options": {
				Type: schema.TypeSet,
				Elem: propertyDomainEncryptionAtRestOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"node_to_node_encryption_options": {
				Type: schema.TypeSet,
				Elem: propertyDomainNodeToNodeEncryptionOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"snapshot_options": {
				Type: schema.TypeSet,
				Elem: propertyDomainSnapshotOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": misc.PropertyTags(),
			"vpc_options": {
				Type: schema.TypeSet,
				Elem: propertyDomainVPCOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceElasticsearchDomainExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceElasticsearchDomainRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticsearchDomainType, ResourceElasticsearchDomain(), data, meta)
}

func resourceElasticsearchDomainCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(elasticsearchDomainType, ResourceElasticsearchDomain(), data, meta)
}

func resourceElasticsearchDomainUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(elasticsearchDomainType, ResourceElasticsearchDomain(), data, meta)
}

func resourceElasticsearchDomainDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(elasticsearchDomainType, data, meta)
}

func resourceElasticsearchDomainCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(elasticsearchDomainType, data, meta)
}
