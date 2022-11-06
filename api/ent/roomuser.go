// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"simple_sns_api/ent/room"
	"simple_sns_api/ent/roomuser"
	"simple_sns_api/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// RoomUser is the model entity for the RoomUser schema.
type RoomUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// RoomID holds the value of the "room_id" field.
	RoomID uuid.UUID `json:"room_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomUserQuery when eager-loading is set.
	Edges RoomUserEdges `json:"edges"`
}

// RoomUserEdges holds the relations/edges for other nodes in the graph.
type RoomUserEdges struct {
	// Room holds the value of the room edge.
	Room *Room `json:"room,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RoomOrErr returns the Room value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomUserEdges) RoomOrErr() (*Room, error) {
	if e.loadedTypes[0] {
		if e.Room == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: room.Label}
		}
		return e.Room, nil
	}
	return nil, &NotLoadedError{edge: "room"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomUserEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoomUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case roomuser.FieldID, roomuser.FieldUserID:
			values[i] = new(sql.NullInt64)
		case roomuser.FieldCreatedAt, roomuser.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case roomuser.FieldRoomID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RoomUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoomUser fields.
func (ru *RoomUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case roomuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ru.ID = int(value.Int64)
		case roomuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ru.CreatedAt = value.Time
			}
		case roomuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ru.UpdatedAt = value.Time
			}
		case roomuser.FieldRoomID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field room_id", values[i])
			} else if value != nil {
				ru.RoomID = *value
			}
		case roomuser.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ru.UserID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryRoom queries the "room" edge of the RoomUser entity.
func (ru *RoomUser) QueryRoom() *RoomQuery {
	return (&RoomUserClient{config: ru.config}).QueryRoom(ru)
}

// QueryUser queries the "user" edge of the RoomUser entity.
func (ru *RoomUser) QueryUser() *UserQuery {
	return (&RoomUserClient{config: ru.config}).QueryUser(ru)
}

// Update returns a builder for updating this RoomUser.
// Note that you need to call RoomUser.Unwrap() before calling this method if this RoomUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (ru *RoomUser) Update() *RoomUserUpdateOne {
	return (&RoomUserClient{config: ru.config}).UpdateOne(ru)
}

// Unwrap unwraps the RoomUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ru *RoomUser) Unwrap() *RoomUser {
	_tx, ok := ru.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoomUser is not a transactional entity")
	}
	ru.config.driver = _tx.drv
	return ru
}

// String implements the fmt.Stringer.
func (ru *RoomUser) String() string {
	var builder strings.Builder
	builder.WriteString("RoomUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ru.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ru.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ru.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("room_id=")
	builder.WriteString(fmt.Sprintf("%v", ru.RoomID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ru.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// RoomUsers is a parsable slice of RoomUser.
type RoomUsers []*RoomUser

func (ru RoomUsers) config(cfg config) {
	for _i := range ru {
		ru[_i].config = cfg
	}
}
