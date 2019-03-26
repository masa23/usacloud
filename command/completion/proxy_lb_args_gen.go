// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBListCompleteArgs(ctx command.Context, params *params.ListProxyLBParam, cur, prev, commandName string) {

}

func ProxyLBCreateCompleteArgs(ctx command.Context, params *params.CreateProxyLBParam, cur, prev, commandName string) {

}

func ProxyLBReadCompleteArgs(ctx command.Context, params *params.ReadProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBUpdateCompleteArgs(ctx command.Context, params *params.UpdateProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBDeleteCompleteArgs(ctx command.Context, params *params.DeleteProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBPlanChangeCompleteArgs(ctx command.Context, params *params.PlanChangeProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBBindPortInfoCompleteArgs(ctx command.Context, params *params.BindPortInfoProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBBindPortAddCompleteArgs(ctx command.Context, params *params.BindPortAddProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBBindPortUpdateCompleteArgs(ctx command.Context, params *params.BindPortUpdateProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBBindPortDeleteCompleteArgs(ctx command.Context, params *params.BindPortDeleteProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBServerInfoCompleteArgs(ctx command.Context, params *params.ServerInfoProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBServerAddCompleteArgs(ctx command.Context, params *params.ServerAddProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBServerUpdateCompleteArgs(ctx command.Context, params *params.ServerUpdateProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBServerDeleteCompleteArgs(ctx command.Context, params *params.ServerDeleteProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBCertificateInfoCompleteArgs(ctx command.Context, params *params.CertificateInfoProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBCertificateAddCompleteArgs(ctx command.Context, params *params.CertificateAddProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBCertificateUpdateCompleteArgs(ctx command.Context, params *params.CertificateUpdateProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBCertificateDeleteCompleteArgs(ctx command.Context, params *params.CertificateDeleteProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ProxyLBMonitorCompleteArgs(ctx command.Context, params *params.MonitorProxyLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProxyLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceProxyLBItems {
		fmt.Println(res.CommonServiceProxyLBItems[i].ID)
		var target interface{} = &res.CommonServiceProxyLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
