package localFileReader

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

type fileReader struct {
filePath string
buffer   []byte
}

func NewFileReader(filePath string, bufferSize int) fileReader {
	buf := make([]byte, bufferSize)

	return fileReader{
		filePath: filePath,
		buffer:   buf,
	}
}

func ReadFile(log *logrus.Logger, rdr fileReader) (string, error) {
	f, err := os.Open(rdr.filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s: %w", rdr.filePath, err)
	}
	defer f.Close()

	var bldr strings.Builder

	for {
		numBytes, err := f.Read(rdr.buffer)

		bldr.Write(rdr.buffer)

		log.Debugf("read %d bytes: %s", numBytes, string(rdr.buffer))

		if err == io.EOF {
			log.Debug("end of file")
			break
		}

		if err != nil {
			return bldr.String(), fmt.Errorf("failed to read file: %w", err)
		}
	}

	return bldr.String(), nil
}
