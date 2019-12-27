// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html

package cloud9

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloud9EnvironmentEC2Type string = "AWS::Cloud9::EnvironmentEC2"

func DatasourceCloud9EnvironmentEC2() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCloud9EnvironmentEC2Read,
		
		Schema: map[string]*schema.Schema{
			"repositories": {
				Type: schema.TypeList,
				Elem: propertyEnvironmentEC2Repository(),
				Optional: true,
			},
			"owner_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"automatic_stop_time_minutes": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
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

func datasourceCloud9EnvironmentEC2Read(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloud9EnvironmentEC2Type, DatasourceCloud9EnvironmentEC2(), data, meta)
}
