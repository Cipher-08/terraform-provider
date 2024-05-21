// resource_yourprovider.go
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceYourProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceYourProviderCreate,
		Read:   resourceYourProviderRead,
		Update: resourceYourProviderUpdate,
		Delete: resourceYourProviderDelete,

		Schema: map[string]*schema.Schema{
			"example": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceYourProviderCreate(d *schema.ResourceData, m interface{}) error {
	// Implementation here
	return nil
}

func resourceYourProviderRead(d *schema.ResourceData, m interface{}) error {
	// Implementation here
	return nil
}

func resourceYourProviderUpdate(d *schema.ResourceData, m interface{}) error {
	// Implementation here
	return nil
}

func resourceYourProviderDelete(d *schema.ResourceData, m interface{}) error {
	// Implementation here
	return nil
}
