package funcs

import (
	"errors"
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBAcmeSetting(ctx command.Context, params *params.AcmeSettingProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBAcmeSetting is failed: %s", e)
	}

	// validate params
	if !params.Disable && !params.AcceptTos {
		return errors.New("--accept-tos=true is required to enable Let's Encrypt setting")
	}

	// set params
	if params.Disable {
		p.Settings.ProxyLB.LetsEncrypt.Enabled = false
		p.Settings.ProxyLB.LetsEncrypt.CommonName = ""
	} else if params.AcceptTos {
		p.Settings.ProxyLB.LetsEncrypt.Enabled = true
		p.Settings.ProxyLB.LetsEncrypt.CommonName = params.CommonName
	}

	// call manipurate functions
	res, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("ProxyLBAcmeSetting is failed: %s", err)
	}
	return ctx.GetOutput().Print(&res.Settings.ProxyLB.LetsEncrypt)
}
