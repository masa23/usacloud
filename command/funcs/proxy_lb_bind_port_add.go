package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBBindPortAdd(ctx command.Context, params *params.BindPortAddProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBBindPortAdd is failed: %s", e)
	}

	// check duplicate
	for _, bp := range p.Settings.ProxyLB.BindPorts {
		if bp.ProxyMode == params.Mode && bp.Port == params.Port {
			return fmt.Errorf("BindPort %s:%d is already exists", bp.ProxyMode, bp.Port)
		}
	}

	p.AddBindPort(params.Mode, params.Port, params.RedirectToHttps, params.SupportHttp2, []*sacloud.ProxyLBResponseHeader{})
	p, e = api.UpdateSetting(params.Id, p)
	if e != nil {
		return fmt.Errorf("ProxyLBBindPortAdd is failed: %s", e)
	}

	var list []interface{}
	for i := range p.Settings.ProxyLB.BindPorts {
		list = append(list, &p.Settings.ProxyLB.BindPorts[i])
	}
	return ctx.GetOutput().Print(list...)

}
