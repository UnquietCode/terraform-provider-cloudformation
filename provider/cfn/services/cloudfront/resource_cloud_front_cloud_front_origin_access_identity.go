// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

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
			"cloud_front_origin_access_identity_config": {
				Type: schema.TypeList,
				Elem: propertyCloudFrontOriginAccessIdentityConfig(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceCloudFrontCloudFrontOriginAccessIdentityCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFront::CloudFrontOriginAccessIdentity", data, meta)
}

func resourceCloudFrontCloudFrontOriginAccessIdentityRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFront::CloudFrontOriginAccessIdentity", data, meta)
}

func resourceCloudFrontCloudFrontOriginAccessIdentityUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFront::CloudFrontOriginAccessIdentity", data, meta)
}

func resourceCloudFrontCloudFrontOriginAccessIdentityDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFront::CloudFrontOriginAccessIdentity", data, meta)
}