// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func BridgeCreateCompleteFlags(ctx Context, params *CreateBridgeParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "description", "desc":
		comp = define.Resources["Bridge"].Commands["create"].Params["description"].CompleteFunc
	case "name":
		comp = define.Resources["Bridge"].Commands["create"].Params["name"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
