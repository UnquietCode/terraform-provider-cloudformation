// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html

package sns

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sNSTopicType string = "AWS::SNS::Topic"

func DatasourceSNSTopic() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSNSTopicRead,
		
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"kms_master_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"subscription": {
				Type: schema.TypeList,
				Elem: propertyTopicSubscription(),
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"topic_name": {
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

func datasourceSNSTopicRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sNSTopicType, DatasourceSNSTopic(), data, meta)
}
