# Golang Review

## Ch. 11 Exercises for 2021_06_09 Meeting

Consider the following example.
  
  a. There is a bug in the code.
  * Write a unit test for the `ReadFile` function.
  * Find and fix the bug. Make sure your unit test(s) include a test case for the bug you found.

  b. Is the `ReadFile` function, as written, following the advice to "Use these interfaces [`io.Reader`, `io.ReadCloser`, etc.] to specify what your functions expect to do with the data"? If not, refactor the `ReadFile` function to change this. How do you think the refactored result compares to the original?

        import (
            "fmt"
            "io"
            "os"
            "strings"

            "github.com/sirupsen/logrus"
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
