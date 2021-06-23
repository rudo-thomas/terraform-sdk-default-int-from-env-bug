package bug

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"test_res": resourceRes(),
		},
	}
}

const (
	defaultVarF = "BUG_DEFAULT_F"
	defaultVarI = "BUG_DEFAULT_I"
)

func resourceRes() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceResCreate,
		ReadContext:   schema.NoopContext,
		UpdateContext: schema.NoopContext,
		DeleteContext: schema.NoopContext,
		Schema: map[string]*schema.Schema{
			"f": &schema.Schema{
				Type:        schema.TypeFloat,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(defaultVarF, nil),
			},
			"i": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(defaultVarI, nil),
			},
		},
		UseJSONNumber: true,
	}
}

func resourceResCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("id")
	return nil
}

func MustSetDefaultsFI(f, i string) {
	if err := os.Setenv(defaultVarF, f); err != nil {
		panic(err)
	}
	if err := os.Setenv(defaultVarI, i); err != nil {
		panic(err)
	}
}
