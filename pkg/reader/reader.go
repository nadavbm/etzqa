package reader

import (
	"io/ioutil"
	"os"

	"github.com/nadavbm/zlog"
	"go.uber.org/zap"
)

// Reader reads files from the file system
type Reader struct {
	logger *zlog.Logger
}

// NewReader creates an instance of a reader
func NewReader(logger *zlog.Logger) *Reader {
	return &Reader{
		logger: logger,
	}
}

// ReadFile get a json or yaml file and return byte slice
func (r *Reader) ReadFile(file string) ([]byte, error) {
	osFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := osFile.Close(); err != nil {
			r.logger.Error("failed to close json file", zap.Error(err))
		}
	}()

	return ioutil.ReadAll(osFile)
}
