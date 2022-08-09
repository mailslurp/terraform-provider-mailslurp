package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/mailslurp/terraform-provider-mailslurp/mailslurp"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mailslurp.Provider})
}
