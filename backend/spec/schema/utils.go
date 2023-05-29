package schema

import (
	"entgo.io/ent/dialect"
	"fmt"
)

const (
	LenAesKey       = 43
	LenAppID        = 25
	LenDate         = 10
	LenDesc         = 255
	LenJSON         = 1530
	LenKey          = 4096
	LenMobile       = 11
	LenName         = 100
	LenNormal       = 50
	LenOrder        = 24
	LenPhone        = 16
	LenURL          = 255
	LenUUID         = 32
	LenTwenty       = 20
	LenFolderName   = 30
	LenFolderNumber = 8

	TransactionTypePay    = "payment"
	TransactionTypeRefund = "refund"
)

var (
	DatetimeSchema = map[string]string{dialect.MySQL: "datetime"}
	DateSchema     = map[string]string{dialect.MySQL: "date"}
	DecimalSchema  = map[string]string{dialect.MySQL: "decimal(12,2)"}
	GeoSchema      = map[string]string{dialect.MySQL: "decimal(9,6)"}

	VarcharFolderName = map[string]string{dialect.MySQL: fmt.Sprintf("varchar(%d)", LenFolderName)}
	VarcharDesc       = map[string]string{dialect.MySQL: fmt.Sprintf("varchar(%d)", 500)}
)
