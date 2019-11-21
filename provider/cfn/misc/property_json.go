package misc

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func PropertyJSON(required bool, extras...string) *schema.Schema {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 15 {
		return &schema.Schema{ Type: schema.TypeString, Optional: true }
	}
	
	return &schema.Schema{
		Type: schema.TypeMap,
    Elem: PropertyJSON(required, strconv.Itoa(int(count + 1))),
    Optional: !required,
    Required: required,
	}
}
