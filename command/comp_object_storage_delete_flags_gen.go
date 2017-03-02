// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func ObjectStorageDeleteCompleteFlags(ctx Context, params *DeleteObjectStorageParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "access-key":
		comp = define.Resources["ObjectStorage"].Commands["delete"].Params["access-key"].CompleteFunc
	case "bucket":
		comp = define.Resources["ObjectStorage"].Commands["delete"].Params["bucket"].CompleteFunc
	case "recursive", "r":
		comp = define.Resources["ObjectStorage"].Commands["delete"].Params["recursive"].CompleteFunc
	case "secret-key":
		comp = define.Resources["ObjectStorage"].Commands["delete"].Params["secret-key"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
