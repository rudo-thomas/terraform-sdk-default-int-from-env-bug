package bug

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestBug(t *testing.T) {
	cfg := `resource "test_res" "r" {}`
	addr := "test_res.r"
	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"test": func() (*schema.Provider, error) { return Provider(), nil },
		},
		Steps: []resource.TestStep{{
			PreConfig: func() { MustSetDefaultsFI("1.23", "") },
			Config:    cfg,
			Check:     resource.TestCheckResourceAttr(addr, "f", "1.23"),
		}, {
			// The bug is in this step. "Error: Attribute must be a whole number, got 1"
			// Tested with:
			//  - github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1
			//  - terraform-1.0.0
			PreConfig: func() { MustSetDefaultsFI("", "1") },
			Config:    cfg,
			Check:     resource.TestCheckResourceAttr(addr, "i", "1"),
		}},
	})
}
