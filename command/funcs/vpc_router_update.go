package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterUpdate(ctx command.Context, params *params.UpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterUpdate is failed: %s", e)
	}

	// set params

	if ctx.IsSet("description") {
		p.SetDescription(params.Description)
	}
	if ctx.IsSet("icon-id") {
		p.SetIconByID(params.IconId)
	}
	if ctx.IsSet("name") {
		p.SetName(params.Name)
	}
	if ctx.IsSet("tags") {
		p.SetTags(params.Tags)
	}
	if ctx.IsSet("syslog-host") {
		p.Settings.Router.SyslogHost = params.SyslogHost
	}
	if ctx.IsSet("internet-connection") {
		p.Settings.Router.InternetConnection = &sacloud.VPCRouterInternetConnection{
			Enabled: "False",
		}
		if params.InternetConnection {
			p.Settings.Router.InternetConnection.Enabled = "True"
		}
	}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterUpdate is failed: %s", err)
	}
	if _, err := api.Config(params.Id); err != nil {
		return fmt.Errorf("VPCRouterUpdate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)
}
