// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-dashboard.html

package cloudwatch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloudWatchDashboardType string = "AWS::CloudWatch::Dashboard"

func ResourceCloudWatchDashboard() *schema.Resource {
	return &schema.Resource{
		Read: resourceCloudWatchDashboardRead,
		Create: resourceCloudWatchDashboardCreate,
		Update: resourceCloudWatchDashboardUpdate,
		Delete: resourceCloudWatchDashboardDelete,
		CustomizeDiff: resourceCloudWatchDashboardCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"dashboard_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"dashboard_body": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCloudWatchDashboardRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloudWatchDashboardType, ResourceCloudWatchDashboard(), data, meta)
}

func resourceCloudWatchDashboardCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cloudWatchDashboardType, ResourceCloudWatchDashboard(), data, meta)
}

func resourceCloudWatchDashboardUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cloudWatchDashboardType, ResourceCloudWatchDashboard(), data, meta)
}

func resourceCloudWatchDashboardDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cloudWatchDashboardType, data, meta)
}

func resourceCloudWatchDashboardCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cloudWatchDashboardType, data, meta)
}
