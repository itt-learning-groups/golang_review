# Ch. 11 Exercises

* Create an original example of "the way to properly wrap types that implement `io.Reader`". Does your example demonstrate how "implementations of `io.Reader` and `io.Writer` are often chained together in a decorator pattern"?

    Ans: `CustomFileTypeReader` is an `io.Reader` that can read files with content such as "key1:val1;key2:val2;key3:val3". Its `Read` method uses another `io.Reader` (an *os.File in this example) and replaces each ":" with a tab, each ";" with a newline.

        package main

        import (
            "fmt"
            "github.com/sirupsen/logrus"
            "io"
            "os"
            "strings"
        )

        const fileName = "test/customFileType/test.custom"

        func main() {

            log := logrus.New()
            log.SetLevel(logrus.DebugLevel)

            f, err := os.Open(fileName)
            if err != nil {
                log.Fatal("failed to open file")
            }
            defer f.Close()

            rdr := CustomFileTypeReader{
                baseReader: f,
                log:        log,
            }
            buf := make([]byte, 4)
            var contents []byte

            for {
                n, err := rdr.Read(buf)

                contents = append(contents, buf[:n]...)

                if err == io.EOF {
                    log.Debug("end of file")
                    break
                }

                if err != nil {
                    log.Error("failed to read from file")
                }
            }

            log.Infof("final read: \n%s", contents)

        }

        type CustomFileTypeReader struct {
            baseReader io.Reader
            log *logrus.Logger
        }

        func (c CustomFileTypeReader) Read(p []byte) (int, error) {
            numBytes, err := c.baseReader.Read(p)
            if err != nil {
                return 0, err
            }

            asStr := string(p)
            asStr = strings.ReplaceAll(asStr, ":", fmt.Sprintf("\t"))
            asStr = strings.ReplaceAll(asStr, ";", fmt.Sprintf("\n"))

            asBytes := []byte(asStr)
            for i := range asBytes {
                p[i] = asBytes[i]
            }

            return numBytes, nil
        }
