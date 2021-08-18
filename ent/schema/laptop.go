package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Laptop holds the schema definition for the Laptop entity.
type Laptop struct {
	ent.Schema
}

// Annotations of the Laptop.
func (Laptop) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tbl_laptop"},
	}
}

// Fields of the Laptop.
func (Laptop) Fields() []ent.Field {
	return []ent.Field{
		field.Text("brand").
			NotEmpty(),
		field.Text("model").
			NotEmpty(),
		field.Text("cpu").
			NotEmpty(),
		field.Text("ram").
			NotEmpty(),
		field.Text("harddisk"),
	}
}

// Edges of the Laptop.
func (Laptop) Edges() []ent.Edge {
	return nil
}
