// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProductDiskList(ctx command.Context, params *params.ListProductDiskParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetProductDiskAPI()

	finder.SetEmpty()

	if !command.IsEmpty(params.Name) {
		for _, v := range params.Name {
			finder.SetFilterBy("Name", v)
		}
	}
	if !command.IsEmpty(params.Id) {
		for _, v := range params.Id {
			finder.SetFilterMultiBy("ID", v)
		}
	}
	if !command.IsEmpty(params.From) {
		finder.SetOffset(params.From)
	}
	if !command.IsEmpty(params.Max) {
		finder.SetLimit(params.Max)
	}
	if !command.IsEmpty(params.Sort) {
		for _, v := range params.Sort {
			setSortBy(finder, v)
		}
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("ProductDiskList is failed: %s", err)
	}

	list := []interface{}{}
	for i := range res.DiskPlans {

		list = append(list, &res.DiskPlans[i])
	}
	return ctx.GetOutput().Print(list...)

}
