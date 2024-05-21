package main

import (
    "context"

    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func providerFunctionEcho() *schema.ProviderFunction {
    return &schema.ProviderFunction{
        Schema: map[string]*schema.Schema{
            "input": {
                Type:     schema.TypeString,
                Required: true,
            },
        },
        Call: func(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
            input := d.Get("input").(string)
            d.Set("output", input)
            return nil
        },
    }
}

