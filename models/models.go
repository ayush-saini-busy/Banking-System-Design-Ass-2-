package models

// Models.go contains the models for GORM as well constants for enums
import (
	"time"
)

// Bank represents a banking org
type Bank struct {
	bankID      int    `gorm:"primaryKey;autoIncrement"`
	name        string `gorm:"type:varchar(255);not null"`
	code        string `gorm:"default:CURRENT_TIMESTAMP"`
	established time.Time
	branches    []Branch `gorm:"foreignKey:bankID"`
}

// Branch represents a brank of a bank
type Branch struct {
	branchID  int        `gorm:"primaryKey;autoIncrement"`
	bankID    int        `gorm:"not null"`
	name      string     `gorm:"type:varchar(255);not null"`
	address   string     `gorm:"type:text;not null"`
	ifscCode  string     `gorm:"type:varchar(11);not null;unique"`
	phone     string     `gorm:"type:varchar(10);unique"`
	bank      Bank       `gorm:"foreignKey:bankID"`
	accounts  []Account  `gorm:"foreignKey:branchID"`
	customers []Customer `gorm:"foreignKey:branchID"`
}

// Customer entity represents a bank customer
type Customer struct {
	customerID int       `gorm:"primaryKey;autoIncrement"`
	name       string    `gorm:"type:varchar(255)"`
	email      string    `gorm:"type:varchar(255)"`
	phone      string    `gorm:"type:varchar(10)"`
	address    string    `gorm:"type:text"`
	createdAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	accounts   []Account `gorm:"foreignKey:customerID"`
	loans      []Loan    `gorm:"foreignKey:customerID"`
}

// Custom type defined to allow only two types of status- "Active" and "Closed"
type AccountStatus string

const (
	AccountStatusActive AccountStatus = "ACTIVE"
	AccountStatusClosed AccountStatus = "CLOSED"
)

// Custome type defined to allow only two types of accounts
type AccountType string

const (
	AccountTypeSavings AccountType = "SAVINGS"
	AccountTypeCurrent AccountType = "CURRENT"
)

// Account represents a customer's account
type Account struct {
	accountID    int           `gorm:"primaryKey"`
	customerID   int           `gorm:"not null"`
	branchID     int           `gorm:"not null"`
	accountNo    string        `gorm:"type:varchar(50);unique;not null"`
	balance      float64       `gorm:"type:decimal(14,2);default:0.00"`
	account_type AccountType   `gorm:"type:varchar(7);default:'SAVINGS'"`
	status       AccountStatus `gorm:"type:varchar(6);default:'ACTIVE'"`
	createdAt    time.Time     `gorm:"default:CURRENT_TIMESTAMP"`
	customer     Customer      `gorm:"foriegnKey:customerID"`
	branch       Branch        `gorm:"foreignKey:branchID"`
	transactions []Transaction `gorm:"foreignKey:accountID"`
}

// Custom type that represents the status of a loan
type LoanStatus string

const (
	LoanActive LoanStatus = "ACTIVE"
	LoanClosed LoanStatus = "CLOSED"
)

// Loan entity represents the loan availed by a customer
type Loan struct {
	LoanID int
}

// Transaction entity represents the transaction performed by a customer
type Transaction struct {
}

type LoanRepayment struct {
}
