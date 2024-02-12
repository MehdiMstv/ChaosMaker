package flagadmin

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

func (g *FlagsGenerator) GetGenerator() map[string]table.Generator {
	return map[string]table.Generator{
		"flags": g.GetTable,
	}
}
