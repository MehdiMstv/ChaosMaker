package chaosadmin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

type ChaosGenerator struct {
}

func (g *ChaosGenerator) GetTable(ctx *context.Context) table.Table {

	chaosTable := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := chaosTable.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int4).
		FieldFilterable().FieldSortable()
	info.AddField("Service", "service", db.Text).
		FieldFilterable().FieldSortable()
	info.AddField("Status", "status", db.Text).
		FieldFilterable().FieldSortable()

	info.SetTable("chaoses").SetTitle("Chaos").SetDescription("Chaos")

	formList := chaosTable.GetForm()
	formList.AddField("ID", "id", db.Integer, form.Default).FieldHide()
	formList.AddField("Service Name", "name", db.Name, form.SelectSingle)
	formList.AddField("Status", "status", db.Text, form.Default).FieldDefault("starting").FieldHide()

	return chaosTable
}
