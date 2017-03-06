// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
)

func ProductDiskReadCompleteArgs(ctx Context, params *ReadProductDiskParam, cur, prev, commandName string) {

	if !GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProductDiskAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}
	for i := range res.DiskPlans {
		fmt.Println(res.DiskPlans[i].ID)
	}

}