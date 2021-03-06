// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func DatabaseListCompleteArgs(ctx command.Context, params *params.ListDatabaseParam, cur, prev, commandName string) {

}

func DatabaseCreateCompleteArgs(ctx command.Context, params *params.CreateDatabaseParam, cur, prev, commandName string) {

}

func DatabaseReadCompleteArgs(ctx command.Context, params *params.ReadDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseUpdateCompleteArgs(ctx command.Context, params *params.UpdateDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseDeleteCompleteArgs(ctx command.Context, params *params.DeleteDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseBootCompleteArgs(ctx command.Context, params *params.BootDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseShutdownCompleteArgs(ctx command.Context, params *params.ShutdownDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseShutdownForceCompleteArgs(ctx command.Context, params *params.ShutdownForceDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseResetCompleteArgs(ctx command.Context, params *params.ResetDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseWaitForBootCompleteArgs(ctx command.Context, params *params.WaitForBootDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseWaitForDownCompleteArgs(ctx command.Context, params *params.WaitForDownDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseBackupInfoCompleteArgs(ctx command.Context, params *params.BackupInfoDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseBackupCreateCompleteArgs(ctx command.Context, params *params.BackupCreateDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseBackupRestoreCompleteArgs(ctx command.Context, params *params.BackupRestoreDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseBackupLockCompleteArgs(ctx command.Context, params *params.BackupLockDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseBackupUnlockCompleteArgs(ctx command.Context, params *params.BackupUnlockDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseBackupRemoveCompleteArgs(ctx command.Context, params *params.BackupRemoveDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseCloneCompleteArgs(ctx command.Context, params *params.CloneDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseReplicaCreateCompleteArgs(ctx command.Context, params *params.ReplicaCreateDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseMonitorCpuCompleteArgs(ctx command.Context, params *params.MonitorCpuDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseMonitorMemoryCompleteArgs(ctx command.Context, params *params.MonitorMemoryDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseMonitorNicCompleteArgs(ctx command.Context, params *params.MonitorNicDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseMonitorSystemDiskCompleteArgs(ctx command.Context, params *params.MonitorSystemDiskDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseMonitorBackupDiskCompleteArgs(ctx command.Context, params *params.MonitorBackupDiskDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseMonitorSystemDiskSizeCompleteArgs(ctx command.Context, params *params.MonitorSystemDiskSizeDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseMonitorBackupDiskSizeCompleteArgs(ctx command.Context, params *params.MonitorBackupDiskSizeDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DatabaseLogsCompleteArgs(ctx command.Context, params *params.LogsDatabaseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDatabaseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Databases {
		fmt.Println(res.Databases[i].ID)
		var target interface{} = &res.Databases[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
