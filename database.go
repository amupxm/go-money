package money

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
)

func (c CAD) Value() (driver.Value, error) {
	return fmt.Sprintf("%d.%d", c.dollar, c.cents), nil
}

func (c *CAD) Scan(src interface{}) (err error) {
	var str string
	v := reflect.ValueOf(src)
	defer func() {
		if errCase := recover(); errCase != nil {
			err = errors.New("invalid money") // to prevent uninitialized panic
		}
	}()
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Slice:
		// what type of slice?
		if v.Type().Elem().Kind() == reflect.String {
			str = v.Interface().(string)
		}
	case reflect.String:
		str = v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		str = fmt.Sprintf("%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		str = fmt.Sprintf("%d", v.Uint())
	default:
		return fmt.Errorf("unsupported type: %s", v.Kind())
	}
	m, err := ParseCAD(str)
	if err != nil {
		return err
	}
	c.dollar = m.dollar
	c.cents = m.cents

	return err
}
