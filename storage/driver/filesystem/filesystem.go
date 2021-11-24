
package filesystem

import (
	"fmt"
	"os"
	"path"
)

type Filesystem struct {
	dataDirectory string
	extension string
}



func New(directory, extension string ) *Filesystem{
	if extension[0:1] != "." {
		extension = "." + extension
	}

	return &Filesystem{directory, extension}
}

//Read gets the file (tableName) on disk and returns a byte array of data stored in the file or an error.
func (f *Filesystem)Read(tableName string) (bytes []byte, err error) {
	bytes, err = os.ReadFile(f.tablePath(tableName))
	if err != nil {
		return nil,fmt.Errorf("unable to read file: %v", err)
	}
	return bytes,nil
}
//Write stores data to a file (tableName) on disk and returns an error status if any.
func (f *Filesystem)Write(tableName string, data []byte) (err error) {
	err = os.WriteFile(f.tablePath(tableName),data,0644)
	if err != nil {
		return fmt.Errorf("error writing data into %s: %w", tableName, err)
	}
	return nil
}

func (f *Filesystem)tablePath(tableName string) (fullPath string){

	return path.Join(f.dataDirectory, tableName + f.extension)
}
