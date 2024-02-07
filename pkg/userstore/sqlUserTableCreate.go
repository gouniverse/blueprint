package userstore

import (
	"github.com/gouniverse/sb"
)

// SQLCreateTable returns a SQL string for creating the country table
func (st *Store) sqlUserTableCreate() string {
	sql := sb.NewBuilder(sb.DatabaseDriverName(st.db)).
		Table(st.userTableName).
		Column(sb.Column{
			Name:       "id",
			Type:       "string",
			PrimaryKey: true,
			Length:     40,
		}).
		Column(sb.Column{
			Name:   "status",
			Type:   "string",
			Length: 40,
		}).
		Column(sb.Column{
			Name:   "first_name",
			Type:   "string",
			Length: 50,
		}).
		Column(sb.Column{
			Name:   "middle_names",
			Type:   "string",
			Length: 50,
		}).
		Column(sb.Column{
			Name:   "last_name",
			Type:   "string",
			Length: 50,
		}).
		Column(sb.Column{
			Name:   "business_name",
			Type:   "string",
			Length: 100,
		}).
		Column(sb.Column{
			Name:   "phone",
			Type:   "string",
			Length: 20,
		}).
		Column(sb.Column{
			Name:   "email",
			Type:   "string",
			Length: 100,
		}).
		Column(sb.Column{
			Name:   "password",
			Type:   "string",
			Length: 255,
		}).
		Column(sb.Column{
			Name:   "role",
			Type:   "string",
			Length: 50,
		}).
		Column(sb.Column{
			Name:   "profile_image_url",
			Type:   "string",
			Length: 255,
		}).
		Column(sb.Column{
			Name: "metas",
			Type: "text",
		}).
		Column(sb.Column{
			Name: "created_at",
			Type: "datetime",
		}).
		Column(sb.Column{
			Name: "updated_at",
			Type: "datetime",
		}).
		Column(sb.Column{
			Name: "deleted_at",
			Type: "datetime",
		}).
		CreateIfNotExists()

	return sql
}
