package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/mirshodNasilloyev/simplebank-go/db/utils"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		//check currency supported
		return utils.IsSupportedCurrency(currency)
	}
	return false
}
