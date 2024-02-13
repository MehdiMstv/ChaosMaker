package serviceadmin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

type ServiceGenerator struct{}

func (g *ServiceGenerator) GetTable(ctx *context.Context) table.Table {

	serviceTable := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := serviceTable.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int4).
		FieldFilterable().FieldSortable()
	info.AddField("Name", "name", db.Text).
		FieldFilterable().FieldSortable()
	info.AddField("Address", "address", db.Text).FieldFilterable().FieldSortable()
	info.AddField("Staging Address", "staging_address", db.Text).
		FieldFilterable().FieldSortable()
	info.AddField("Chaos Address", "chaos_address", db.Text).
		FieldFilterable().FieldSortable()

	info.SetTable("services").SetTitle("Services").SetDescription("Services")

	formList := serviceTable.GetForm()
	formList.AddField("Name", "name", db.Text, form.Text).FieldMust()
	formList.AddField("Address", "address", db.Text, form.Text).FieldMust()
	formList.AddField("Staging Address", "staging_address", db.Text, form.Text)
	formList.AddField("Chaos Address", "chaos_address", db.Text, form.Text)
	formList.SetTable("services").SetTitle("Services").SetDescription("Services")

	return serviceTable
}
