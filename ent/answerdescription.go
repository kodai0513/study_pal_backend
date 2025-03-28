// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"study-pal-backend/ent/answerdescription"
	"study-pal-backend/ent/problem"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// AnswerDescription is the model entity for the AnswerDescription schema.
type AnswerDescription struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ProblemID holds the value of the "problem_id" field.
	ProblemID uuid.UUID `json:"problem_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AnswerDescriptionQuery when eager-loading is set.
	Edges        AnswerDescriptionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AnswerDescriptionEdges holds the relations/edges for other nodes in the graph.
type AnswerDescriptionEdges struct {
	// Problem holds the value of the problem edge.
	Problem *Problem `json:"problem,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProblemOrErr returns the Problem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AnswerDescriptionEdges) ProblemOrErr() (*Problem, error) {
	if e.Problem != nil {
		return e.Problem, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: problem.Label}
	}
	return nil, &NotLoadedError{edge: "problem"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AnswerDescription) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case answerdescription.FieldName:
			values[i] = new(sql.NullString)
		case answerdescription.FieldCreatedAt, answerdescription.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case answerdescription.FieldID, answerdescription.FieldProblemID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AnswerDescription fields.
func (ad *AnswerDescription) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case answerdescription.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ad.ID = *value
			}
		case answerdescription.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ad.CreatedAt = value.Time
			}
		case answerdescription.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ad.UpdatedAt = value.Time
			}
		case answerdescription.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ad.Name = value.String
			}
		case answerdescription.FieldProblemID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field problem_id", values[i])
			} else if value != nil {
				ad.ProblemID = *value
			}
		default:
			ad.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AnswerDescription.
// This includes values selected through modifiers, order, etc.
func (ad *AnswerDescription) Value(name string) (ent.Value, error) {
	return ad.selectValues.Get(name)
}

// QueryProblem queries the "problem" edge of the AnswerDescription entity.
func (ad *AnswerDescription) QueryProblem() *ProblemQuery {
	return NewAnswerDescriptionClient(ad.config).QueryProblem(ad)
}

// Update returns a builder for updating this AnswerDescription.
// Note that you need to call AnswerDescription.Unwrap() before calling this method if this AnswerDescription
// was returned from a transaction, and the transaction was committed or rolled back.
func (ad *AnswerDescription) Update() *AnswerDescriptionUpdateOne {
	return NewAnswerDescriptionClient(ad.config).UpdateOne(ad)
}

// Unwrap unwraps the AnswerDescription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ad *AnswerDescription) Unwrap() *AnswerDescription {
	_tx, ok := ad.config.driver.(*txDriver)
	if !ok {
		panic("ent: AnswerDescription is not a transactional entity")
	}
	ad.config.driver = _tx.drv
	return ad
}

// String implements the fmt.Stringer.
func (ad *AnswerDescription) String() string {
	var builder strings.Builder
	builder.WriteString("AnswerDescription(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ad.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ad.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ad.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ad.Name)
	builder.WriteString(", ")
	builder.WriteString("problem_id=")
	builder.WriteString(fmt.Sprintf("%v", ad.ProblemID))
	builder.WriteByte(')')
	return builder.String()
}

// AnswerDescriptions is a parsable slice of AnswerDescription.
type AnswerDescriptions []*AnswerDescription
