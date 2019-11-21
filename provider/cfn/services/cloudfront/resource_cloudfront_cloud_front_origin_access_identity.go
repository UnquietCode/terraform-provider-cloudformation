// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html

package cloudfront

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloudFrontCloudFrontOriginAccessIdentityType string = "AWS::CloudFront::CloudFrontOriginAccessIdentity"

func ResourceCloudFrontCloudFrontOriginAccessIdentity() *schema.Resource {
	return &schema.Resource{
		Read: resourceCloudFrontCloudFrontOriginAccessIdentityRead,
		Create: resourceCloudFrontCloudFrontOriginAccessIdentityCreate,
		Update: resourceCloudFrontCloudFrontOriginAccessIdentityUpdate,
		Delete: resourceCloudFrontCloudFrontOriginAccessIdentityDelete,
		CustomizeDiff: resourceCloudFrontCloudFrontOriginAccessIdentityCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"cloud_front_origin_access_identity_config": {
				Type: schema.TypeSet,
				Elem: propertyCloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfig(),
				Required: true,
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

func resourceCloudFrontCloudFrontOriginAccessIdentityRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloudFrontCloudFrontOriginAccessIdentityType, ResourceCloudFrontCloudFrontOriginAccessIdentity(), data, meta)
}

func resourceCloudFrontCloudFrontOriginAccessIdentityCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cloudFrontCloudFrontOriginAccessIdentityType, ResourceCloudFrontCloudFrontOriginAccessIdentity(), data, meta)
}

func resourceCloudFrontCloudFrontOriginAccessIdentityUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cloudFrontCloudFrontOriginAccessIdentityType, ResourceCloudFrontCloudFrontOriginAccessIdentity(), data, meta)
}

func resourceCloudFrontCloudFrontOriginAccessIdentityDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cloudFrontCloudFrontOriginAccessIdentityType, data, meta)
}

func resourceCloudFrontCloudFrontOriginAccessIdentityCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cloudFrontCloudFrontOriginAccessIdentityType, data, meta)
}
