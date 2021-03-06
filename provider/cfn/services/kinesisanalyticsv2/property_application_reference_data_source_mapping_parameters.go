// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-mappingparameters.html

package kinesisanalyticsv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationReferenceDataSourceMappingParameters(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
	    if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
	        count = i
	    }
	}
	
	if count >= 5 {
	    return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"json_mapping_parameters": {
				Type: schema.TypeList,
				Elem: propertyApplicationReferenceDataSourceJSONMappingParameters(),
				Optional: true,
				MaxItems: 1,
			},
			"csv_mapping_parameters": {
				Type: schema.TypeList,
				Elem: propertyApplicationReferenceDataSourceCSVMappingParameters(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
