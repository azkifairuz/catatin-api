// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type CategoryType string

const (
	CategoryTypeExpense CategoryType = "expense"
	CategoryTypeIncome  CategoryType = "income"
)

func (e *CategoryType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CategoryType(s)
	case string:
		*e = CategoryType(s)
	default:
		return fmt.Errorf("unsupported scan type for CategoryType: %T", src)
	}
	return nil
}

type NullCategoryType struct {
	CategoryType CategoryType
	Valid        bool // Valid is true if CategoryType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCategoryType) Scan(value interface{}) error {
	if value == nil {
		ns.CategoryType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CategoryType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCategoryType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CategoryType), nil
}

type IntervalType string

const (
	IntervalTypeDaily   IntervalType = "daily"
	IntervalTypeWeekly  IntervalType = "weekly"
	IntervalTypeMonthly IntervalType = "monthly"
)

func (e *IntervalType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = IntervalType(s)
	case string:
		*e = IntervalType(s)
	default:
		return fmt.Errorf("unsupported scan type for IntervalType: %T", src)
	}
	return nil
}

type NullIntervalType struct {
	IntervalType IntervalType
	Valid        bool // Valid is true if IntervalType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullIntervalType) Scan(value interface{}) error {
	if value == nil {
		ns.IntervalType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.IntervalType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullIntervalType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.IntervalType), nil
}

type Account struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type BudgetPlan struct {
	ID         int64
	Name       pgtype.Text
	Amount     pgtype.Float8
	CategoryID pgtype.Int8
	AccountsID int64
	Interval   NullIntervalType
}

type Cashflow struct {
	ID           int64
	Amount       pgtype.Float8
	Name         pgtype.Text
	WalletID     int64
	BudgetPlanID int64
	CategoryID   int64
}

type Category struct {
	ID         int64
	Name       string
	Category   CategoryType
	AccountsID pgtype.Int8
}

type Wallet struct {
	ID         int64
	Name       string
	Balance    float64
	AccountsID pgtype.Int8
}
