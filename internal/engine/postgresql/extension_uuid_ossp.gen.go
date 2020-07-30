package postgresql

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
	"github.com/kyleconroy/sqlc/internal/sql/catalog"
)

func genUUIDOSSP() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = []*catalog.Function{
		{
			Name:       "uuid_generate_v1",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "uuid"},
		},
		{
			Name:       "uuid_generate_v1mc",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "uuid"},
		},
		{
			Name: "uuid_generate_v3",
			Args: []*catalog.Argument{
				{
					Name: "namespace",
					Type: &ast.TypeName{Name: "uuid"},
				},
				{
					Name: "name",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "uuid"},
		},
		{
			Name:       "uuid_generate_v4",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "uuid"},
		},
		{
			Name: "uuid_generate_v5",
			Args: []*catalog.Argument{
				{
					Name: "namespace",
					Type: &ast.TypeName{Name: "uuid"},
				},
				{
					Name: "name",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "uuid"},
		},
		{
			Name:       "uuid_nil",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "uuid"},
		},
		{
			Name:       "uuid_ns_dns",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "uuid"},
		},
		{
			Name:       "uuid_ns_oid",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "uuid"},
		},
		{
			Name:       "uuid_ns_url",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "uuid"},
		},
		{
			Name:       "uuid_ns_x500",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "uuid"},
		},
	}
	return s
}
