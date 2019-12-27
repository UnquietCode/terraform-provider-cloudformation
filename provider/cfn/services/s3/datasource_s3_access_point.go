// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html

package s3

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const s3AccessPointType string = "AWS::S3::AccessPoint"

func DatasourceS3AccessPoint() *schema.Resource {
	return &schema.Resource{
		Read: datasourceS3AccessPointRead,
		
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"bucket": {
				Type: schema.TypeString,
				Required: true,
			},
			"vpc_configuration": {
				Type: schema.TypeList,
				Elem: propertyAccessPointVpcConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"public_access_block_configuration": {
				Type: schema.TypeList,
				Elem: propertyAccessPointPublicAccessBlockConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"policy": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"policy_status": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"network_origin": {
				Type: schema.TypeString,
				Optional: true,
			},
			"creation_date": {
				Type: schema.TypeString,
				Optional: true,
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

func datasourceS3AccessPointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(s3AccessPointType, DatasourceS3AccessPoint(), data, meta)
}
