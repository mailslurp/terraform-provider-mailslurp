package mailslurp

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceMailSlurpDomain() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,

				ForceNew: true,
			},
		},
	}
}
