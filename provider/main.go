package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/unquietcode/terraform-cfn-provider/cfn"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cfn.Provider})
}