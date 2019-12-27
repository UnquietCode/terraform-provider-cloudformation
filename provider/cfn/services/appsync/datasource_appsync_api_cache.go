// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html

package appsync

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appSyncApiCacheType string = "AWS::AppSync::ApiCache"

func DatasourceAppSyncApiCache() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAppSyncApiCacheRead,
		
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"transit_encryption_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"at_rest_encryption_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"api_caching_behavior": {
				Type: schema.TypeString,
				Required: true,
			},
			"ttl": {
				Type: schema.TypeFloat,
				Required: true,
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

func datasourceAppSyncApiCacheRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appSyncApiCacheType, DatasourceAppSyncApiCache(), data, meta)
}
