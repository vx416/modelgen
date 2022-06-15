package dbhelper

type Helper interface {
	IsString(typ string) bool
	IsText(typ string) bool
	IsInteger(typ string) bool
	IsSmallInterger(typ string) bool
	IsFloat(typ string) bool
	IsTimestamp(typ string) bool
}
