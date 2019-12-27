// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html

package gamelift

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const gameLiftGameSessionQueueType string = "AWS::GameLift::GameSessionQueue"

func DatasourceGameLiftGameSessionQueue() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGameLiftGameSessionQueueRead,
		
		Schema: map[string]*schema.Schema{
			"timeout_in_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"player_latency_policies": {
				Type: schema.TypeList,
				Elem: propertyGameSessionQueuePlayerLatencyPolicy(),
				Optional: true,
			},
			"destinations": {
				Type: schema.TypeList,
				Elem: propertyGameSessionQueueDestination(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
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

func datasourceGameLiftGameSessionQueueRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(gameLiftGameSessionQueueType, DatasourceGameLiftGameSessionQueue(), data, meta)
}
