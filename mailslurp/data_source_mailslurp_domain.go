package mailslurp

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMailSlurpDomain() *schema.Resource {

	mailslurpSchema := resourceMailSlurpDomain()

	return &schema.Resource{
		Schema: mailslurpSchema.Schema,
	}
}
