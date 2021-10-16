package driver

import (
	"fmt"
)

type Storage interface {
	Write(string, []byte) error
	Read(string) ([]byte, error)
}

type Driver struct {
	storageLayer Storage
}

func New(storageLayer Storage) *Driver{
	return &Driver{storageLayer: storageLayer}
}

func (D* Driver)Insert( tableName string, data []byte) (err error) {
	err = D.storageLayer.Write(tableName,data)
	if err != nil {
		return fmt.Errorf("running database insert data into %s: %w", tableName, err)
	}
	return nil
}

func (D* Driver)Read( table string) (bytes []byte,err error) {
	bytes, err = D.storageLayer.Read(table)
	if err != nil {
		return nil,fmt.Errorf("running database read %s: %w", table, err)
	}
	return bytes,nil
}




