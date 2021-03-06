package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func RegionResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             regionListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: regionListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"read": {
			Type:          schema.CommandManipulateIDOnly,
			Params:        regionReadParam(),
			IncludeFields: regionDetailIncludes(),
			ExcludeFields: regionDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		DefaultCommand:   "list",
		ResourceCategory: CategoryInformation,
	}
}

func regionListParam() map[string]*schema.Schema {
	return CommonListParam
}

func regionListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{Name: "Description"},
		{
			Name:    "NameServers",
			Sources: []string{"NameServers.0", "NameServers.1"},
			Format:  "%s,%s",
		},
	}
}

func regionDetailIncludes() []string {
	return []string{}
}

func regionDetailExcludes() []string {
	return []string{}
}

func regionReadParam() map[string]*schema.Schema {
	id := getParamResourceShortID("resource ID", 3)
	id.Hidden = true
	return map[string]*schema.Schema{
		"id": id,
	}
}
