// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html

package elasticsearch

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticsearchDomainType string = "AWS::Elasticsearch::Domain"

func DatasourceElasticsearchDomain() *schema.Resource {
	return &schema.Resource{
		Read: datasourceElasticsearchDomainRead,
		
		Schema: map[string]*schema.Schema{
			"access_policies": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"advanced_options": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"cognito_options": {
				Type: schema.TypeList,
				Elem: propertyDomainCognitoOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"domain_name": {
				Type: schema.TypeString,
				Optional: true,
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
			},
			"encryption_at_rest_options": {
				Type: schema.TypeList,
				Elem: propertyDomainEncryptionAtRestOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"log_publishing_options": {
				Type: schema.TypeMap,
				Elem: propertyDomainLogPublishingOption(),
				Optional: true,
			},
			"node_to_node_encryption_options": {
				Type: schema.TypeList,
				Elem: propertyDomainNodeToNodeEncryptionOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"snapshot_options": {
				Type: schema.TypeList,
				Elem: propertyDomainSnapshotOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": misc.PropertyTags(),
			"vpc_options": {
				Type: schema.TypeList,
				Elem: propertyDomainVPCOptions(),
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

func datasourceElasticsearchDomainRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticsearchDomainType, DatasourceElasticsearchDomain(), data, meta)
}
