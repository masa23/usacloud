// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func DNSListCompleteFlags(ctx command.Context, params *params.ListDNSParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["DNS"].Commands["list"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["DNS"].Commands["list"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags", "selector":
		param := define.Resources["DNS"].Commands["list"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "from", "offset":
		param := define.Resources["DNS"].Commands["list"].BuildedParams().Get("from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "max", "limit":
		param := define.Resources["DNS"].Commands["list"].BuildedParams().Get("max")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sort":
		param := define.Resources["DNS"].Commands["list"].BuildedParams().Get("sort")
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

func DNSRecordInfoCompleteFlags(ctx command.Context, params *params.RecordInfoDNSParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["DNS"].Commands["record-info"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "type":
		param := define.Resources["DNS"].Commands["record-info"].BuildedParams().Get("type")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["DNS"].Commands["record-info"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["DNS"].Commands["record-info"].BuildedParams().Get("id")
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

func DNSCreateCompleteFlags(ctx command.Context, params *params.CreateDNSParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["DNS"].Commands["create"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["DNS"].Commands["create"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["DNS"].Commands["create"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "icon-id":
		param := define.Resources["DNS"].Commands["create"].BuildedParams().Get("icon-id")
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

func DNSRecordAddCompleteFlags(ctx command.Context, params *params.RecordAddDNSParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["DNS"].Commands["record-add"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "type":
		param := define.Resources["DNS"].Commands["record-add"].BuildedParams().Get("type")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "value":
		param := define.Resources["DNS"].Commands["record-add"].BuildedParams().Get("value")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "ttl":
		param := define.Resources["DNS"].Commands["record-add"].BuildedParams().Get("ttl")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "mx-priority":
		param := define.Resources["DNS"].Commands["record-add"].BuildedParams().Get("mx-priority")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "srv-priority":
		param := define.Resources["DNS"].Commands["record-add"].BuildedParams().Get("srv-priority")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "srv-weight":
		param := define.Resources["DNS"].Commands["record-add"].BuildedParams().Get("srv-weight")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "srv-port":
		param := define.Resources["DNS"].Commands["record-add"].BuildedParams().Get("srv-port")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "srv-target":
		param := define.Resources["DNS"].Commands["record-add"].BuildedParams().Get("srv-target")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["DNS"].Commands["record-add"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["DNS"].Commands["record-add"].BuildedParams().Get("id")
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

func DNSReadCompleteFlags(ctx command.Context, params *params.ReadDNSParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["DNS"].Commands["read"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["DNS"].Commands["read"].BuildedParams().Get("id")
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

func DNSRecordUpdateCompleteFlags(ctx command.Context, params *params.RecordUpdateDNSParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["DNS"].Commands["record-update"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "name":
		param := define.Resources["DNS"].Commands["record-update"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "type":
		param := define.Resources["DNS"].Commands["record-update"].BuildedParams().Get("type")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "value":
		param := define.Resources["DNS"].Commands["record-update"].BuildedParams().Get("value")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "ttl":
		param := define.Resources["DNS"].Commands["record-update"].BuildedParams().Get("ttl")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "mx-priority":
		param := define.Resources["DNS"].Commands["record-update"].BuildedParams().Get("mx-priority")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "srv-priority":
		param := define.Resources["DNS"].Commands["record-update"].BuildedParams().Get("srv-priority")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "srv-weight":
		param := define.Resources["DNS"].Commands["record-update"].BuildedParams().Get("srv-weight")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "srv-port":
		param := define.Resources["DNS"].Commands["record-update"].BuildedParams().Get("srv-port")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "srv-target":
		param := define.Resources["DNS"].Commands["record-update"].BuildedParams().Get("srv-target")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["DNS"].Commands["record-update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["DNS"].Commands["record-update"].BuildedParams().Get("id")
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

func DNSRecordDeleteCompleteFlags(ctx command.Context, params *params.RecordDeleteDNSParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["DNS"].Commands["record-delete"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["DNS"].Commands["record-delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["DNS"].Commands["record-delete"].BuildedParams().Get("id")
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

func DNSUpdateCompleteFlags(ctx command.Context, params *params.UpdateDNSParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["DNS"].Commands["update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["DNS"].Commands["update"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["DNS"].Commands["update"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "icon-id":
		param := define.Resources["DNS"].Commands["update"].BuildedParams().Get("icon-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["DNS"].Commands["update"].BuildedParams().Get("id")
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

func DNSDeleteCompleteFlags(ctx command.Context, params *params.DeleteDNSParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["DNS"].Commands["delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["DNS"].Commands["delete"].BuildedParams().Get("id")
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
