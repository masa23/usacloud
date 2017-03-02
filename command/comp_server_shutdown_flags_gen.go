// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func ServerShutdownCompleteFlags(ctx Context, params *ShutdownServerParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "async":
		comp = define.Resources["Server"].Commands["shutdown"].Params["async"].CompleteFunc
	case "force":
		comp = define.Resources["Server"].Commands["shutdown"].Params["force"].CompleteFunc
	case "id":
		comp = define.Resources["Server"].Commands["shutdown"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
