// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceElasticsearchDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceElasticsearchDomainCreate,
		Read:   resourceElasticsearchDomainRead,
		Update: resourceElasticsearchDomainUpdate,
		Delete: resourceElasticsearchDomainDelete,

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
				ForceNew: true,
			},
			"ebs_options": {
				Type: schema.TypeList,
				Elem: propertyDomainEBSOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"elasticsearch_cluster_config": {
				Type: schema.TypeList,
				Elem: propertyDomainElasticsearchClusterConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"elasticsearch_version": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"encryption_at_rest_options": {
				Type: schema.TypeList,
				Elem: propertyDomainEncryptionAtRestOptions(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"node_to_node_encryption_options": {
				Type: schema.TypeList,
				Elem: propertyDomainNodeToNodeEncryptionOptions(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"snapshot_options": {
				Type: schema.TypeList,
				Elem: propertyDomainSnapshotOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"vpc_options": {
				Type: schema.TypeList,
				Elem: propertyDomainVPCOptions(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceElasticsearchDomainCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Elasticsearch::Domain", data, meta)
}

func resourceElasticsearchDomainRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Elasticsearch::Domain", data, meta)
}

func resourceElasticsearchDomainUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Elasticsearch::Domain", data, meta)
}

func resourceElasticsearchDomainDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Elasticsearch::Domain", data, meta)
}