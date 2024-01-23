package chaosadmin

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

func (g *ChaosGenerator) GetGenerator() map[string]table.Generator {
	return map[string]table.Generator{
		"chaos": g.GetTable,
	}
}
