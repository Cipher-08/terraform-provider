// functions.go
package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func providerFunctionEcho() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"input": {
				Type:     schema.TypeString,
				Required: true,
			},
			"output": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			input := d.Get("input").(string)
			d.Set("output", input)
			d.SetId("echo")
			return nil
		},
	}
}
