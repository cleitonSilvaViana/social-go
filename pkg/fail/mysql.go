package fail

import (
	"net/http"

	"github.com/go-sql-driver/mysql"
)

type MySqlError struct {
	mysql.MySQLError
	Code int
	Erro string
}

func (me *MySqlError) setStatusCodeError() {
	switch me.Number {
	case 1000, 1062, 1002, 1452:
		me.Code = http.StatusBadRequest
		me.Erro = me.Message
	default:
		me.Code = http.StatusInternalServerError
		me.Erro = "internal server error"
	}
}

func NewMySqlError(err error) *ResponseError {
	dbError := MySqlError{}

	if e, ok := err.(*mysql.MySQLError); ok {
		dbError.Number = e.Number
		dbError.SQLState = e.SQLState
		dbError.Message = e.Message

		dbError.setStatusCodeError()
	} else {
		dbError.Code = http.StatusInternalServerError
		dbError.Message = "internal server error"
	}

	// criar um log de erro aqui....

	return &ResponseError{
		StatusCode: dbError.Code,
		Message:    dbError.Message,
	}
}
