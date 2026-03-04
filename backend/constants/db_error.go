package constants

type MysqlError struct {
	ErrCode int
}

func (m MysqlError) Error() string {
	return "mysql error with code: " + string(m.ErrCode)
}

var (
	ErrUserNotFound = MysqlError{ErrCode: 0}
)
