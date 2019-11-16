package misc

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func PropertyTags(extras...string) *schema.Schema {
	return &schema.Schema{
		Type: schema.TypeMap,
    Elem: &schema.Schema{Type: schema.TypeString},
    Optional: true,
	}
}
