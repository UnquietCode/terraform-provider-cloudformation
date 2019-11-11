// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html

package cloudfront

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCloudFrontCloudFrontOriginAccessIdentity() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFrontCloudFrontOriginAccessIdentityCreate,
		Read:   resourceCloudFrontCloudFrontOriginAccessIdentityRead,
		Update: resourceCloudFrontCloudFrontOriginAccessIdentityUpdate,
		Delete: resourceCloudFrontCloudFrontOriginAccessIdentityDelete,

		Schema: map[string]*schema.Schema{
			"s3_canonical_user_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"cloud_front_origin_access_identity_config": {
				Type: schema.TypeList,
				Elem: propertyCloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfig(),
				Required: true,
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

func resourceCloudFrontCloudFrontOriginAccessIdentityCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFront::CloudFrontOriginAccessIdentity", ResourceCloudFrontCloudFrontOriginAccessIdentity(), data, meta)
}

func resourceCloudFrontCloudFrontOriginAccessIdentityRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFront::CloudFrontOriginAccessIdentity", ResourceCloudFrontCloudFrontOriginAccessIdentity(), data, meta)
}

func resourceCloudFrontCloudFrontOriginAccessIdentityUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFront::CloudFrontOriginAccessIdentity", ResourceCloudFrontCloudFrontOriginAccessIdentity(), data, meta)
}

func resourceCloudFrontCloudFrontOriginAccessIdentityDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFront::CloudFrontOriginAccessIdentity", data, meta)
}