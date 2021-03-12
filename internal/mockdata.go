package internal

import (
	"encoding/json"
	"fmt"
	"reflect"

	pgorm "github.com/go-pg/pg/v10/orm"
	"github.com/go-pg/pg/v10/types"
)

var (
	dstErr    = fmt.Errorf("dst type is not ptr")
	dstNilErr = fmt.Errorf("dst is nil ptr")
	paramErr  = fmt.Errorf("param type is not ptr")
)

func MockData(dst interface{}, param interface{}) error {
	if dst == nil {
		return dstNilErr
	}

	table, err := pgorm.NewModel(dst)
	if err != nil {
		return dstErr
	}
	err = table.Init()
	if err != nil {
		return err
	}

	model := reflect.ValueOf(param)
	if model.Kind() == reflect.Ptr {
		model = model.Elem()
	}

	return mock(table, model)
}

func mock(table pgorm.Model, model reflect.Value) error {
	switch model.Kind() {
	case reflect.Slice:
		for i := 0; i < model.Cap(); i++ {
			v := model.Index(i)
			if v.Kind() != reflect.Struct {
				return paramErr
			}
			scanner := table.NextColumnScanner()
			err := mockScanner(scanner, v)
			if err != nil {
				return err
			}
			err = table.AddColumnScanner(scanner)
			if err != nil {
				return err
			}
		}
	case reflect.Struct:
		scanner := table.NextColumnScanner()
		err := mockScanner(scanner, model)
		if err != nil {
			return err
		}
		return table.AddColumnScanner(scanner)
	default:
		return paramErr
	}
	return nil
}

func mockScanner(scanner pgorm.ColumnScanner, model reflect.Value) error {
	modelV := model.Type()
	for i := 0; i < modelV.NumField(); i++ {
		field := modelV.Field(i)
		column := Underscore(field.Name)
		if !model.Field(i).CanInterface() {
			continue
		}
		typ := model.Field(i).Interface()
		data, _ := json.Marshal(typ)
		if len(data) > 0 && data[0] == '"' {
			data = data[1 : len(data)-1]
		}
		rd := NewBytesReader(data)
		err := scanner.ScanColumn(types.ColumnInfo{
			Index:    int16(i),
			Name:     column,
			DataType: int32(field.Type.Kind()),
		}, rd, rd.Buffered())
		if err != nil {
			return fmt.Errorf("invalid mock data: %v", err)
		}
	}
	return nil
}
