// data_source_yourprovider.go
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceYourProvider() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceYourProviderRead,

        Schema: map[string]*schema.Schema{
            "example": {
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func dataSourceYourProviderRead(d *schema.ResourceData, m interface{}) error {
    // Implementation here
    return nil
}
