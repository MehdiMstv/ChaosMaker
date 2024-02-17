package flagadmin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

type FlagsGenerator struct {
	Conn db.Connection
}

func (g *FlagsGenerator) GetTable(ctx *context.Context) table.Table {

	flagsTable := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := flagsTable.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int4).
		FieldFilterable().FieldSortable()
	info.AddField("Name", "name", db.Text).
		FieldFilterable().FieldSortable()
	info.AddField("Service Name", "service_name", db.Text).
		FieldFilterable().FieldSortable()
	info.AddField("Type", "type", db.Integer).
		FieldFilterable().FieldSortable().FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "String"
		}
		if model.Value == "1" {
			return "Boolean"
		}
		if model.Value == "2" {
			return "Integer"
		}
		return "Unknown"
	})
	info.AddField("Value", "value", db.Text).
		FieldFilterable().FieldSortable()
	info.AddField("Staging Value", "staging_value", db.Text).
		FieldFilterable().FieldSortable()

	info.SetTable("flags").SetTitle("Flags").SetDescription("Flags")

	formList := flagsTable.GetForm()
	formList.AddField("ID", "id", db.Integer, form.Default).FieldHide()
	formList.AddField("Name", "name", db.Name, form.Text).FieldMust()

	formList.AddField("Service", "service_name", db.Text, form.SelectSingle).FieldOptions(getServices(g.Conn))
	formList.AddField("Type", "type", db.Integer, form.SelectSingle).FieldOptions([]types.FieldOption{
		{Text: "String", Value: "0"},
		{Text: "Boolean", Value: "1"},
		{Text: "Integer", Value: "2"},
	})
	formList.AddField("Value", "value", db.Text, form.Text)
	formList.AddField("Staging Value", "staging_value", db.Text, form.Text)

	formList.SetTable("flags").SetTitle("Flags").SetDescription("Flags")

	return flagsTable
}

func getServices(conn db.Connection) []types.FieldOption {
	query, err := conn.Query("Select name from services")
	if err != nil {
		return nil
	}
	var services []types.FieldOption
	for _, v := range query {
		services = append(services, types.FieldOption{
			Text:  v["name"].(string),
			Value: v["name"].(string),
		})
	}
	return services
}
