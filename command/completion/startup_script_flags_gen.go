// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func StartupScriptListCompleteFlags(ctx command.Context, params *params.ListStartupScriptParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		comp = define.Resources["StartupScript"].Commands["list"].Params["name"].CompleteFunc
	case "id":
		comp = define.Resources["StartupScript"].Commands["list"].Params["id"].CompleteFunc
	case "scope":
		comp = define.Resources["StartupScript"].Commands["list"].Params["scope"].CompleteFunc
	case "tags":
		comp = define.Resources["StartupScript"].Commands["list"].Params["tags"].CompleteFunc
	case "from", "offset":
		comp = define.Resources["StartupScript"].Commands["list"].Params["from"].CompleteFunc
	case "max", "limit":
		comp = define.Resources["StartupScript"].Commands["list"].Params["max"].CompleteFunc
	case "sort":
		comp = define.Resources["StartupScript"].Commands["list"].Params["sort"].CompleteFunc
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func StartupScriptCreateCompleteFlags(ctx command.Context, params *params.CreateStartupScriptParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "script", "note":
		comp = define.Resources["StartupScript"].Commands["create"].Params["script"].CompleteFunc
	case "script-content", "note-content":
		comp = define.Resources["StartupScript"].Commands["create"].Params["script-content"].CompleteFunc
	case "name":
		comp = define.Resources["StartupScript"].Commands["create"].Params["name"].CompleteFunc
	case "tags":
		comp = define.Resources["StartupScript"].Commands["create"].Params["tags"].CompleteFunc
	case "icon-id":
		comp = define.Resources["StartupScript"].Commands["create"].Params["icon-id"].CompleteFunc
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func StartupScriptReadCompleteFlags(ctx command.Context, params *params.ReadStartupScriptParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["StartupScript"].Commands["read"].Params["id"].CompleteFunc
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func StartupScriptUpdateCompleteFlags(ctx command.Context, params *params.UpdateStartupScriptParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "script", "note":
		comp = define.Resources["StartupScript"].Commands["update"].Params["script"].CompleteFunc
	case "script-content", "note-content":
		comp = define.Resources["StartupScript"].Commands["update"].Params["script-content"].CompleteFunc
	case "name":
		comp = define.Resources["StartupScript"].Commands["update"].Params["name"].CompleteFunc
	case "tags":
		comp = define.Resources["StartupScript"].Commands["update"].Params["tags"].CompleteFunc
	case "icon-id":
		comp = define.Resources["StartupScript"].Commands["update"].Params["icon-id"].CompleteFunc
	case "id":
		comp = define.Resources["StartupScript"].Commands["update"].Params["id"].CompleteFunc
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func StartupScriptDeleteCompleteFlags(ctx command.Context, params *params.DeleteStartupScriptParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["StartupScript"].Commands["delete"].Params["id"].CompleteFunc
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}