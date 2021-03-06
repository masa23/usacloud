package funcs

import (
	"fmt"
	"io/ioutil"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SSHKeyCreate(ctx command.Context, params *params.CreateSSHKeyParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSSHKeyAPI()
	p := api.New()

	// validate
	if params.PublicKey == "" && params.PublicKeyContent == "" {
		return fmt.Errorf("%q or %q is required", "public-key", "public-key-content")
	}

	// set params
	if params.PublicKey != "" {
		b, err := ioutil.ReadFile(params.PublicKey)
		if err != nil {
			return fmt.Errorf("SSHKeyCreate is failed: %s", err)
		}
		p.PublicKey = string(b)
	}

	if params.PublicKeyContent != "" {
		p.PublicKey = params.PublicKeyContent
	}

	p.SetName(params.Name)
	p.SetDescription(params.Description)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("SSHKeyCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
