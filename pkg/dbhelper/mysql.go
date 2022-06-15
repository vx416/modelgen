package dbhelper

import "strings"

type MySQL struct {
}

// GetStringDatatypes returns the string datatypes for the MySQL database.
func (mysql MySQL) GetStringDatatypes() []string {
	return []string{
		"char",
		"varchar",
		"binary",
		"varbinary",
	}
}

// IsString returns true if the colum is of type string for the MySQL database.
func (mysql MySQL) IsString(typ string) bool {
	return mysql.IsStringInSlice(typ, mysql.GetStringDatatypes())
}

// IsText returns true if colum is of type text for the MySQL database.
func (mysql MySQL) IsText(typ string) bool {
	return mysql.IsStringInSlice(typ, mysql.GetTextDatatypes())
}

// IsInteger returns true if colum is of type integer for the MySQL database.
func (mysql MySQL) IsInteger(typ string) bool {
	return mysql.IsStringInSlice(typ, mysql.GetIntegerDatatypes())
}

// IsInteger returns true if colum is of type integer for the MySQL database.
func (mysql MySQL) IsSmallInterger(typ string) bool {
	return mysql.IsStringInSlice(typ, mysql.GetSmallIntegerDatatypes())
}

// IsFloat returns true if colum is of type float for the MySQL database.
func (mysql MySQL) IsFloat(typ string) bool {
	return mysql.IsStringInSlice(typ, mysql.GetFloatDatatypes())
}

func (mysql MySQL) IsTimestamp(typ string) bool {
	return mysql.IsStringInSlice(typ, mysql.GetTemporalDatatypes())
}

// IsStringInSlice checks if needle (string) is in haystack ([]string).
func (gdb MySQL) IsStringInSlice(needle string, haystack []string) bool {
	for _, s := range haystack {
		if strings.EqualFold(s, needle) {
			return true
		}
	}
	return false
}

// GetTextDatatypes returns the text datatypes for the MySQL database.
func (mysql MySQL) GetTextDatatypes() []string {
	return []string{
		"text",
		"blob",
	}
}

// GetIntegerDatatypes returns the integer datatypes for the MySQL database.
func (mysql MySQL) GetIntegerDatatypes() []string {
	return []string{
		"mediumint",
		"int",
		"bigint",
	}
}

// GetIntegerDatatypes returns the integer datatypes for the MySQL database.
func (mysql MySQL) GetSmallIntegerDatatypes() []string {
	return []string{
		"tinyint",
		"smallint",
	}
}

// GetFloatDatatypes returns the float datatypes for the MySQL database.
func (mysql MySQL) GetFloatDatatypes() []string {
	return []string{
		"numeric",
		"decimal",
		"float",
		"real",
		"double precision",
	}
}

func (mysql MySQL) GetTemporalDatatypes() []string {
	return []string{
		"time",
		"timestamp",
		"date",
		"datetime",
		"year",
	}
}
