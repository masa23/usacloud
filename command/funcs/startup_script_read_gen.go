// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func StartupScriptRead(ctx command.Context, params *params.ReadStartupScriptParam) error {

	client := ctx.GetAPIClient()
	api := client.GetNoteAPI()

	// set params

	// call Read(id)
	res, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("StartupScriptRead is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
