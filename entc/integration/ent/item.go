// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/item"
)

// Item is the model entity for the Item schema.
type Item struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Item) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Item fields.
func (i *Item) assignValues(values ...interface{}) error {
	if m, n := len(values), len(item.Columns); m != n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	i.ID = strconv.FormatInt(value.Int64, 10)
	values = values[1:]
	return nil
}

// Update returns a builder for updating this Item.
// Note that, you need to call Item.Unwrap() before calling this method, if this Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Item) Update() *ItemUpdateOne {
	return (&ItemClient{i.config}).UpdateOne(i)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (i *Item) Unwrap() *Item {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Item is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Item) String() string {
	var builder strings.Builder
	builder.WriteString("Item(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (i *Item) id() int {
	id, _ := strconv.Atoi(i.ID)
	return id
}

// Items is a parsable slice of Item.
type Items []*Item

func (i Items) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
