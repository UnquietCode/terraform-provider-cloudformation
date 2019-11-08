// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ecs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyContainerDefinition() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"cpu": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"depends_on": {
				Type: schema.TypeSet,
				Elem: propertyContainerDependency(),
				Required: false,
				ForceNew: true,
			},
			"disable_networking": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"dns_search_domains": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"dns_servers": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"docker_labels": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"docker_security_options": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"entry_point": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"environment": {
				Type: schema.TypeSet,
				Elem: propertyKeyValuePair(),
				Required: false,
				ForceNew: true,
			},
			"essential": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"extra_hosts": {
				Type: schema.TypeSet,
				Elem: propertyHostEntry(),
				Required: false,
				ForceNew: true,
			},
			"health_check": {
				Type: schema.TypeList,
				Elem: propertyHealthCheck(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"hostname": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"image": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"interactive": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"links": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"linux_parameters": {
				Type: schema.TypeList,
				Elem: propertyLinuxParameters(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"log_configuration": {
				Type: schema.TypeList,
				Elem: propertyLogConfiguration(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"memory": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"memory_reservation": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"mount_points": {
				Type: schema.TypeSet,
				Elem: propertyMountPoint(),
				Required: false,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"port_mappings": {
				Type: schema.TypeSet,
				Elem: propertyPortMapping(),
				Required: false,
				ForceNew: true,
			},
			"privileged": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"pseudo_terminal": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"readonly_root_filesystem": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"repository_credentials": {
				Type: schema.TypeList,
				Elem: propertyRepositoryCredentials(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"resource_requirements": {
				Type: schema.TypeSet,
				Elem: propertyResourceRequirement(),
				Required: false,
				ForceNew: true,
			},
			"secrets": {
				Type: schema.TypeSet,
				Elem: propertySecret(),
				Required: false,
				ForceNew: true,
			},
			"start_timeout": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"stop_timeout": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"system_controls": {
				Type: schema.TypeSet,
				Elem: propertySystemControl(),
				Required: false,
				ForceNew: true,
			},
			"ulimits": {
				Type: schema.TypeSet,
				Elem: propertyUlimit(),
				Required: false,
				ForceNew: true,
			},
			"user": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"volumes_from": {
				Type: schema.TypeSet,
				Elem: propertyVolumeFrom(),
				Required: false,
				ForceNew: true,
			},
			"working_directory": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}