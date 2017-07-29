// TARの解凍定義

package tar

import (
	"archive/tar"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/naga718/golang-practice/ch10/ex02/archive"
)

func init() {
	archive.RegisterFormat("tar", "./", Unzip)
}

func Unzip(tarfile string, dir string) (err error) {
	var file *os.File

	file, err = os.Open(tarfile)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := tar.NewReader(file)

	var header *tar.Header
	for {
		header, err = reader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}

		path := filepath.Join(dir, header.Name)
		err = ioutil.WriteFile(path, buf, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
