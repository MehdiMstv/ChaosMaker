package chaosadmin

import (
	"errors"
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"net/http"
)

type ChaosGenerator struct {
	Conn db.Connection
}

func (g *ChaosGenerator) GetTable(ctx *context.Context) table.Table {

	chaosTable := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := chaosTable.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int4).FieldFilterable().FieldSortable()
	info.AddField("Service", "service_name", db.Text).FieldFilterable().FieldSortable()
	info.AddField("Method", "method", db.Text).FieldFilterable().FieldSortable()
	info.AddField("Status", "status", db.Text).FieldFilterable().FieldSortable()
	info.AddField("Result", "result", db.Text).FieldFilterable().FieldSortable()
	info.AddField("Created At", "created_at", db.Datetime).FieldFilterable().FieldSortable()

	info.SetTable("chaos").SetTitle("Chaos").SetDescription("Chaos").HideEditButton().HideDeleteButton()

	formList := chaosTable.GetForm()
	fieldOptions := getServices(g.Conn)
	formList.AddField("Service Name", "service_name", db.Text, form.SelectSingle).FieldOptions(fieldOptions).FieldMust()
	formList.AddField("Status", "status", db.Text, form.Default).FieldHide().FieldDefault("Unknown")
	formList.AddField("Created At", "created_at", db.Text, form.Default).FieldHide().FieldNow()
	formList.AddField("Method", "method", db.Text, form.Text)

	formList.SetTable("chaos").SetTitle("Chaos").SetDescription("Chaos").SetPostHook(func(values form2.Values) error {
		chaosAddress := getChaosAddress(g.Conn, values.Get("service_name"))
		if chaosAddress == "" {
			return errors.New("empty chaos address")
		}
		_, err := http.Post(fmt.Sprintf("http://%s/start_%s_chaos?id=%v", chaosAddress, values.Get("method"), values.Get("id")), "application/json", nil)
		if err != nil {
			return err
		}
		return nil
	})
	return chaosTable
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

func getChaosAddress(conn db.Connection, name string) string {
	queryText := fmt.Sprintf(`Select chaos_address from services where name='%s'`, name)
	query, err := conn.Query(queryText)
	if err != nil || len(query) != 1 || query[0]["chaos_address"] == nil {
		return ""
	}
	return query[0]["chaos_address"].(string)
}
