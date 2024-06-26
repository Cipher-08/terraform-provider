// provider.go
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"yourprovider_resource": resourceYourProvider(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"yourprovider_data": dataSourceYourProvider(),
			"yourprovider_echo": providerFunctionEcho(),
		},
		Schema: map[string]*schema.Schema{},
	}
}
