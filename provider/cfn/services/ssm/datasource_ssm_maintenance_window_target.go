// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtarget.html

package ssm

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sSMMaintenanceWindowTargetType string = "AWS::SSM::MaintenanceWindowTarget"

func DatasourceSSMMaintenanceWindowTarget() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSSMMaintenanceWindowTargetRead,
		
		Schema: map[string]*schema.Schema{
			"owner_information": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"window_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"targets": {
				Type: schema.TypeList,
				Elem: propertyMaintenanceWindowTargetTargets(),
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

func datasourceSSMMaintenanceWindowTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sSMMaintenanceWindowTargetType, DatasourceSSMMaintenanceWindowTarget(), data, meta)
}
