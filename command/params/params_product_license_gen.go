// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListProductLicenseParam is input parameters for the sacloud API
type ListProductLicenseParam struct {
	Name              []string `json:"name"`
	Id                []int64  `json:"id"`
	From              int      `json:"from"`
	Max               int      `json:"max"`
	Sort              []string `json:"sort"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	QueryFile         string   `json:"query-file"`
}

// NewListProductLicenseParam return new ListProductLicenseParam
func NewListProductLicenseParam() *ListProductLicenseParam {
	return &ListProductLicenseParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ListProductLicenseParam) FillValueToSkeleton() {
	if isEmpty(p.Name) {
		p.Name = []string{""}
	}
	if isEmpty(p.Id) {
		p.Id = []int64{0}
	}
	if isEmpty(p.From) {
		p.From = 0
	}
	if isEmpty(p.Max) {
		p.Max = 0
	}
	if isEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

// Validate checks current values in model
func (p *ListProductLicenseParam) Validate() []error {
	errors := []error{}
	{
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ProductLicense"].Commands["list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListProductLicenseParam) GetResourceDef() *schema.Resource {
	return define.Resources["ProductLicense"]
}

func (p *ListProductLicenseParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListProductLicenseParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListProductLicenseParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListProductLicenseParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListProductLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListProductLicenseParam) SetName(v []string) {
	p.Name = v
}

func (p *ListProductLicenseParam) GetName() []string {
	return p.Name
}
func (p *ListProductLicenseParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListProductLicenseParam) GetId() []int64 {
	return p.Id
}
func (p *ListProductLicenseParam) SetFrom(v int) {
	p.From = v
}

func (p *ListProductLicenseParam) GetFrom() int {
	return p.From
}
func (p *ListProductLicenseParam) SetMax(v int) {
	p.Max = v
}

func (p *ListProductLicenseParam) GetMax() int {
	return p.Max
}
func (p *ListProductLicenseParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListProductLicenseParam) GetSort() []string {
	return p.Sort
}
func (p *ListProductLicenseParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListProductLicenseParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListProductLicenseParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListProductLicenseParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListProductLicenseParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListProductLicenseParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListProductLicenseParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListProductLicenseParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListProductLicenseParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListProductLicenseParam) GetColumn() []string {
	return p.Column
}
func (p *ListProductLicenseParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListProductLicenseParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListProductLicenseParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListProductLicenseParam) GetFormat() string {
	return p.Format
}
func (p *ListProductLicenseParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListProductLicenseParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListProductLicenseParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListProductLicenseParam) GetQuery() string {
	return p.Query
}
func (p *ListProductLicenseParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListProductLicenseParam) GetQueryFile() string {
	return p.QueryFile
}

// ReadProductLicenseParam is input parameters for the sacloud API
type ReadProductLicenseParam struct {
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	QueryFile         string   `json:"query-file"`
	Id                int64    `json:"id"`
}

// NewReadProductLicenseParam return new ReadProductLicenseParam
func NewReadProductLicenseParam() *ReadProductLicenseParam {
	return &ReadProductLicenseParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ReadProductLicenseParam) FillValueToSkeleton() {
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}
	if isEmpty(p.Id) {
		p.Id = 0
	}

}

// Validate checks current values in model
func (p *ReadProductLicenseParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ProductLicense"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadProductLicenseParam) GetResourceDef() *schema.Resource {
	return define.Resources["ProductLicense"]
}

func (p *ReadProductLicenseParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["read"]
}

func (p *ReadProductLicenseParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ReadProductLicenseParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ReadProductLicenseParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ReadProductLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ReadProductLicenseParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *ReadProductLicenseParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *ReadProductLicenseParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadProductLicenseParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadProductLicenseParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadProductLicenseParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadProductLicenseParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadProductLicenseParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadProductLicenseParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadProductLicenseParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadProductLicenseParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadProductLicenseParam) GetColumn() []string {
	return p.Column
}
func (p *ReadProductLicenseParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadProductLicenseParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadProductLicenseParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadProductLicenseParam) GetFormat() string {
	return p.Format
}
func (p *ReadProductLicenseParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadProductLicenseParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadProductLicenseParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadProductLicenseParam) GetQuery() string {
	return p.Query
}
func (p *ReadProductLicenseParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadProductLicenseParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadProductLicenseParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadProductLicenseParam) GetId() int64 {
	return p.Id
}
