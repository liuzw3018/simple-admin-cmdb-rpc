package schema

import "entgo.io/ent"

// Cmdb holds the schema definition for the Cmdb entity.
type Cmdb struct {
	ent.Schema
}

// Fields of the Cmdb.
func (Cmdb) Fields() []ent.Field {
	return nil
}

// Edges of the Cmdb.
func (Cmdb) Edges() []ent.Edge {
	return nil
}
