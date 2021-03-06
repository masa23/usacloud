// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func PrivateHostListCompleteArgs(ctx command.Context, params *params.ListPrivateHostParam, cur, prev, commandName string) {

}

func PrivateHostCreateCompleteArgs(ctx command.Context, params *params.CreatePrivateHostParam, cur, prev, commandName string) {

}

func PrivateHostReadCompleteArgs(ctx command.Context, params *params.ReadPrivateHostParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPrivateHostAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PrivateHosts {
		fmt.Println(res.PrivateHosts[i].ID)
		var target interface{} = &res.PrivateHosts[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func PrivateHostUpdateCompleteArgs(ctx command.Context, params *params.UpdatePrivateHostParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPrivateHostAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PrivateHosts {
		fmt.Println(res.PrivateHosts[i].ID)
		var target interface{} = &res.PrivateHosts[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func PrivateHostDeleteCompleteArgs(ctx command.Context, params *params.DeletePrivateHostParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPrivateHostAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PrivateHosts {
		fmt.Println(res.PrivateHosts[i].ID)
		var target interface{} = &res.PrivateHosts[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func PrivateHostServerInfoCompleteArgs(ctx command.Context, params *params.ServerInfoPrivateHostParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPrivateHostAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PrivateHosts {
		fmt.Println(res.PrivateHosts[i].ID)
		var target interface{} = &res.PrivateHosts[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func PrivateHostServerAddCompleteArgs(ctx command.Context, params *params.ServerAddPrivateHostParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPrivateHostAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PrivateHosts {
		fmt.Println(res.PrivateHosts[i].ID)
		var target interface{} = &res.PrivateHosts[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func PrivateHostServerDeleteCompleteArgs(ctx command.Context, params *params.ServerDeletePrivateHostParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPrivateHostAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.PrivateHosts {
		fmt.Println(res.PrivateHosts[i].ID)
		var target interface{} = &res.PrivateHosts[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
