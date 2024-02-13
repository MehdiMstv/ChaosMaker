package serviceadmin

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

func (g *ServiceGenerator) GetGenerator() map[string]table.Generator {
	return map[string]table.Generator{
		"services": g.GetTable,
	}
}
