// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func MobileGatewayListCompleteFlags(ctx command.Context, params *params.ListMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["MobileGateway"].Commands["list"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["list"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags", "selector":
		param := define.Resources["MobileGateway"].Commands["list"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "from", "offset":
		param := define.Resources["MobileGateway"].Commands["list"].BuildedParams().Get("from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "max", "limit":
		param := define.Resources["MobileGateway"].Commands["list"].BuildedParams().Get("max")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sort":
		param := define.Resources["MobileGateway"].Commands["list"].BuildedParams().Get("sort")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayCreateCompleteFlags(ctx command.Context, params *params.CreateMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "internet-connection":
		param := define.Resources["MobileGateway"].Commands["create"].BuildedParams().Get("internet-connection")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "name":
		param := define.Resources["MobileGateway"].Commands["create"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["MobileGateway"].Commands["create"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["MobileGateway"].Commands["create"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "icon-id":
		param := define.Resources["MobileGateway"].Commands["create"].BuildedParams().Get("icon-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayReadCompleteFlags(ctx command.Context, params *params.ReadMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["MobileGateway"].Commands["read"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["read"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayUpdateCompleteFlags(ctx command.Context, params *params.UpdateMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "internet-connection":
		param := define.Resources["MobileGateway"].Commands["update"].BuildedParams().Get("internet-connection")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "name":
		param := define.Resources["MobileGateway"].Commands["update"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["MobileGateway"].Commands["update"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["MobileGateway"].Commands["update"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "icon-id":
		param := define.Resources["MobileGateway"].Commands["update"].BuildedParams().Get("icon-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayDeleteCompleteFlags(ctx command.Context, params *params.DeleteMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "force", "f":
		param := define.Resources["MobileGateway"].Commands["delete"].BuildedParams().Get("force")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayBootCompleteFlags(ctx command.Context, params *params.BootMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["MobileGateway"].Commands["boot"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["boot"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayShutdownCompleteFlags(ctx command.Context, params *params.ShutdownMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["MobileGateway"].Commands["shutdown"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["shutdown"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayShutdownForceCompleteFlags(ctx command.Context, params *params.ShutdownForceMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["MobileGateway"].Commands["shutdown-force"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["shutdown-force"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayResetCompleteFlags(ctx command.Context, params *params.ResetMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["MobileGateway"].Commands["reset"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["reset"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayWaitForBootCompleteFlags(ctx command.Context, params *params.WaitForBootMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["MobileGateway"].Commands["wait-for-boot"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["wait-for-boot"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayWaitForDownCompleteFlags(ctx command.Context, params *params.WaitForDownMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["MobileGateway"].Commands["wait-for-down"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["wait-for-down"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayInterfaceInfoCompleteFlags(ctx command.Context, params *params.InterfaceInfoMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["MobileGateway"].Commands["interface-info"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["interface-info"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayInterfaceConnectCompleteFlags(ctx command.Context, params *params.InterfaceConnectMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "ipaddress", "ip":
		param := define.Resources["MobileGateway"].Commands["interface-connect"].BuildedParams().Get("ipaddress")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "switch-id":
		param := define.Resources["MobileGateway"].Commands["interface-connect"].BuildedParams().Get("switch-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "nw-masklen", "network-masklen":
		param := define.Resources["MobileGateway"].Commands["interface-connect"].BuildedParams().Get("nw-masklen")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["interface-connect"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["interface-connect"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayInterfaceUpdateCompleteFlags(ctx command.Context, params *params.InterfaceUpdateMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "ipaddress", "ip":
		param := define.Resources["MobileGateway"].Commands["interface-update"].BuildedParams().Get("ipaddress")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "nw-masklen", "network-masklen":
		param := define.Resources["MobileGateway"].Commands["interface-update"].BuildedParams().Get("nw-masklen")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["interface-update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["interface-update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayInterfaceDisconnectCompleteFlags(ctx command.Context, params *params.InterfaceDisconnectMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["MobileGateway"].Commands["interface-disconnect"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["interface-disconnect"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayTrafficControlInfoCompleteFlags(ctx command.Context, params *params.TrafficControlInfoMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["MobileGateway"].Commands["traffic-control-info"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["traffic-control-info"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayTrafficControlEnableCompleteFlags(ctx command.Context, params *params.TrafficControlEnableMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "quota":
		param := define.Resources["MobileGateway"].Commands["traffic-control-enable"].BuildedParams().Get("quota")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "band-width-limit":
		param := define.Resources["MobileGateway"].Commands["traffic-control-enable"].BuildedParams().Get("band-width-limit")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "enable-email":
		param := define.Resources["MobileGateway"].Commands["traffic-control-enable"].BuildedParams().Get("enable-email")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "slack-webhook-url":
		param := define.Resources["MobileGateway"].Commands["traffic-control-enable"].BuildedParams().Get("slack-webhook-url")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "auto-traffic-shaping":
		param := define.Resources["MobileGateway"].Commands["traffic-control-enable"].BuildedParams().Get("auto-traffic-shaping")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["traffic-control-enable"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["traffic-control-enable"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayTrafficControlUpdateCompleteFlags(ctx command.Context, params *params.TrafficControlUpdateMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "quota":
		param := define.Resources["MobileGateway"].Commands["traffic-control-update"].BuildedParams().Get("quota")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "band-width-limit":
		param := define.Resources["MobileGateway"].Commands["traffic-control-update"].BuildedParams().Get("band-width-limit")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "enable-email":
		param := define.Resources["MobileGateway"].Commands["traffic-control-update"].BuildedParams().Get("enable-email")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "slack-webhook-url":
		param := define.Resources["MobileGateway"].Commands["traffic-control-update"].BuildedParams().Get("slack-webhook-url")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "auto-traffic-shaping":
		param := define.Resources["MobileGateway"].Commands["traffic-control-update"].BuildedParams().Get("auto-traffic-shaping")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["traffic-control-update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["traffic-control-update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayTrafficControlDisableCompleteFlags(ctx command.Context, params *params.TrafficControlDisableMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["MobileGateway"].Commands["traffic-control-disable"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["traffic-control-disable"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayStaticRouteInfoCompleteFlags(ctx command.Context, params *params.StaticRouteInfoMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["MobileGateway"].Commands["static-route-info"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["static-route-info"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayStaticRouteAddCompleteFlags(ctx command.Context, params *params.StaticRouteAddMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "prefix":
		param := define.Resources["MobileGateway"].Commands["static-route-add"].BuildedParams().Get("prefix")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "next-hop":
		param := define.Resources["MobileGateway"].Commands["static-route-add"].BuildedParams().Get("next-hop")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["static-route-add"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["static-route-add"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayStaticRouteUpdateCompleteFlags(ctx command.Context, params *params.StaticRouteUpdateMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["MobileGateway"].Commands["static-route-update"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "prefix":
		param := define.Resources["MobileGateway"].Commands["static-route-update"].BuildedParams().Get("prefix")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "next-hop":
		param := define.Resources["MobileGateway"].Commands["static-route-update"].BuildedParams().Get("next-hop")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["static-route-update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["static-route-update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayStaticRouteDeleteCompleteFlags(ctx command.Context, params *params.StaticRouteDeleteMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["MobileGateway"].Commands["static-route-delete"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["static-route-delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["static-route-delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewaySimInfoCompleteFlags(ctx command.Context, params *params.SimInfoMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["MobileGateway"].Commands["sim-info"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["sim-info"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewaySimAddCompleteFlags(ctx command.Context, params *params.SimAddMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "ipaddress", "ip":
		param := define.Resources["MobileGateway"].Commands["sim-add"].BuildedParams().Get("ipaddress")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sim-id":
		param := define.Resources["MobileGateway"].Commands["sim-add"].BuildedParams().Get("sim-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["sim-add"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["sim-add"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewaySimUpdateCompleteFlags(ctx command.Context, params *params.SimUpdateMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "ipaddress", "ip":
		param := define.Resources["MobileGateway"].Commands["sim-update"].BuildedParams().Get("ipaddress")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sim-id":
		param := define.Resources["MobileGateway"].Commands["sim-update"].BuildedParams().Get("sim-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["sim-update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["sim-update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewaySimDeleteCompleteFlags(ctx command.Context, params *params.SimDeleteMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "sim-id":
		param := define.Resources["MobileGateway"].Commands["sim-delete"].BuildedParams().Get("sim-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["sim-delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["sim-delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewaySimRouteInfoCompleteFlags(ctx command.Context, params *params.SimRouteInfoMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["MobileGateway"].Commands["sim-route-info"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["sim-route-info"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewaySimRouteAddCompleteFlags(ctx command.Context, params *params.SimRouteAddMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "prefix":
		param := define.Resources["MobileGateway"].Commands["sim-route-add"].BuildedParams().Get("prefix")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sim":
		param := define.Resources["MobileGateway"].Commands["sim-route-add"].BuildedParams().Get("sim")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["sim-route-add"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["sim-route-add"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewaySimRouteUpdateCompleteFlags(ctx command.Context, params *params.SimRouteUpdateMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["MobileGateway"].Commands["sim-route-update"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "prefix":
		param := define.Resources["MobileGateway"].Commands["sim-route-update"].BuildedParams().Get("prefix")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sim":
		param := define.Resources["MobileGateway"].Commands["sim-route-update"].BuildedParams().Get("sim")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["sim-route-update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["sim-route-update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewaySimRouteDeleteCompleteFlags(ctx command.Context, params *params.SimRouteDeleteMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["MobileGateway"].Commands["sim-route-delete"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["sim-route-delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["sim-route-delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayDnsUpdateCompleteFlags(ctx command.Context, params *params.DnsUpdateMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "dns1":
		param := define.Resources["MobileGateway"].Commands["dns-update"].BuildedParams().Get("dns1")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "dns2":
		param := define.Resources["MobileGateway"].Commands["dns-update"].BuildedParams().Get("dns2")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["dns-update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["dns-update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func MobileGatewayLogsCompleteFlags(ctx command.Context, params *params.LogsMobileGatewayParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "follow", "f":
		param := define.Resources["MobileGateway"].Commands["logs"].BuildedParams().Get("follow")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "refresh-interval":
		param := define.Resources["MobileGateway"].Commands["logs"].BuildedParams().Get("refresh-interval")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["MobileGateway"].Commands["logs"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["MobileGateway"].Commands["logs"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
