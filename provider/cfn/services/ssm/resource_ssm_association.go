// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSSMAssociation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSSMAssociationExists,
		Read:   resourceSSMAssociationRead,
		Create: resourceSSMAssociationCreate,
		Update: resourceSSMAssociationUpdate,
		Delete: resourceSSMAssociationDelete,
		
		Schema: map[string]*schema.Schema{
			"association_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"document_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"output_location": {
				Type: schema.TypeList,
				Elem: propertyAssociationInstanceAssociationOutputLocation(),
				Optional: true,
				MaxItems: 1,
			},
			"parameters": {
				Type: schema.TypeMap,
				Elem: propertyAssociationParameterValues(),
				Optional: true,
			},
			"schedule_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"targets": {
				Type: schema.TypeSet,
				Elem: propertyAssociationTarget(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSSMAssociationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSSMAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::Association", ResourceSSMAssociation(), data, meta)
}

func resourceSSMAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::Association", ResourceSSMAssociation(), data, meta)
}

func resourceSSMAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::Association", ResourceSSMAssociation(), data, meta)
}

func resourceSSMAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::Association", data, meta)
}