// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func SwitchBridgeConnectCompleteFlags(ctx Context, params *BridgeConnectSwitchParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "bridge-id":
		comp = define.Resources["Switch"].Commands["bridge-connect"].Params["bridge-id"].CompleteFunc
	case "id":
		comp = define.Resources["Switch"].Commands["bridge-connect"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}