// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
)

func SSHKeyDeleteCompleteArgs(ctx Context, params *DeleteSSHKeyParam, cur, prev, commandName string) {

	if !GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSSHKeyAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}
	for i := range res.SSHKeys {
		fmt.Println(res.SSHKeys[i].ID)
	}

}