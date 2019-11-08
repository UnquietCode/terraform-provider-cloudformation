// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSSMAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSMAssociationCreate,
		Read:   resourceSSMAssociationRead,
		Update: resourceSSMAssociationUpdate,
		Delete: resourceSSMAssociationDelete,

		Schema: map[string]*schema.Schema{
			"association_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"document_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"instance_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"output_location": {
				Type: schema.TypeList,
				Elem: propertyInstanceAssociationOutputLocation(),
				Required: false,
				MaxItems: 1,
			},
			"parameters": {
				Type: schema.TypeMap,
				Elem: propertyParameterValues(),
				Required: false,
			},
			"schedule_expression": {
				Type: schema.TypeString,
				Required: false,
			},
			"targets": {
				Type: schema.TypeSet,
				Elem: propertyTarget(),
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceSSMAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::Association", data, meta)
}

func resourceSSMAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::Association", data, meta)
}

func resourceSSMAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::Association", data, meta)
}

func resourceSSMAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::Association", data, meta)
}