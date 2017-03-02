// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
)

func AutoBackupUpdateCompleteArgs(ctx Context, params *UpdateAutoBackupParam) {

	if !GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetAutoBackupAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}
	for i := range res.CommonServiceAutoBackupItems {
		fmt.Println(res.CommonServiceAutoBackupItems[i].ID)
	}

}
