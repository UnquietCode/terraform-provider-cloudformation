// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html

package docdb

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const docDBDBInstanceType string = "AWS::DocDB::DBInstance"

func ResourceDocDBDBInstance() *schema.Resource {
	return &schema.Resource{
		Read: resourceDocDBDBInstanceRead,
		Create: resourceDocDBDBInstanceCreate,
		Update: resourceDocDBDBInstanceUpdate,
		Delete: resourceDocDBDBInstanceDelete,
		CustomizeDiff: resourceDocDBDBInstanceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"db_instance_class": {
				Type: schema.TypeString,
				Required: true,
			},
			"db_cluster_identifier": {
				Type: schema.TypeString,
				Required: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Optional: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"auto_minor_version_upgrade": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"db_instance_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceDocDBDBInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(docDBDBInstanceType, ResourceDocDBDBInstance(), data, meta)
}

func resourceDocDBDBInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(docDBDBInstanceType, ResourceDocDBDBInstance(), data, meta)
}

func resourceDocDBDBInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(docDBDBInstanceType, ResourceDocDBDBInstance(), data, meta)
}

func resourceDocDBDBInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(docDBDBInstanceType, data, meta)
}

func resourceDocDBDBInstanceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(docDBDBInstanceType, data, meta)
}
