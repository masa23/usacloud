// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func ServerIsoInsertCompleteFlags(ctx Context, params *IsoInsertServerParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "description", "desc":
		comp = define.Resources["Server"].Commands["iso-insert"].Params["description"].CompleteFunc
	case "icon-id":
		comp = define.Resources["Server"].Commands["iso-insert"].Params["icon-id"].CompleteFunc
	case "id":
		comp = define.Resources["Server"].Commands["iso-insert"].Params["id"].CompleteFunc
	case "iso-file":
		comp = define.Resources["Server"].Commands["iso-insert"].Params["iso-file"].CompleteFunc
	case "iso-image-id":
		comp = define.Resources["Server"].Commands["iso-insert"].Params["iso-image-id"].CompleteFunc
	case "name":
		comp = define.Resources["Server"].Commands["iso-insert"].Params["name"].CompleteFunc
	case "size":
		comp = define.Resources["Server"].Commands["iso-insert"].Params["size"].CompleteFunc
	case "tags":
		comp = define.Resources["Server"].Commands["iso-insert"].Params["tags"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
