# Ch. 11 Exercises

## 1. Consider the following example.
  
### a. There is a bug in the code.

* Write a unit test for the `ReadFile` function.
* Find and fix the bug. Make sure your unit test(s) include a test case for the bug you found.

### b. Is the `ReadFile` function, as written, following the advice to "Use these interfaces [`io.Reader`, `io.ReadCloser`, etc.] to specify what your functions expect to do with the data"? If not, refactor the `ReadFile` function to change this. How do you think the refactored result compares to the original?

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

----------

Ans / example soln(s):

a. When we read less than `len(rdr.buffer)` bytes (i.e. when we get to the end of the file), the `rdr.buffer` in the line `bldr.Write(rdr.buffer)` will contain bytes from the previous reading loop. So we'll get repeated bytes in our output string. Update the line so we only add to the string builder the number of bytes we actually read: `bldr.Write(rdr.buffer[:numBytes])`.

    import (
        "fmt"
        "github.com/sirupsen/logrus"
        "os"
        "testing"
    )

    func Test(t *testing.T) {
        testFileContents := `1  This royal throne of kings, this sceptred isle,
            2  This earth of majesty, this seat of Mars,
            3  This other Eden, demi-paradise,
            4  This fortress built by Nature for her self
            5  Against infection and the hand of war...`

        testFile, err := os.CreateTemp("", "testFile")
        if err != nil {
            t.Fatal(err)
        }
        defer os.Remove(testFile.Name())

        if _, err := testFile.Write([]byte(testFileContents)); err != nil {
            t.Fatal(err)
        }
        if err := testFile.Close(); err != nil {
            t.Fatal(err)
        }

        log := logrus.New()
        bufferSize := 128

        fileContents, err := ReadFile(log, NewFileReader(testFile.Name(), bufferSize))

        fmt.Printf("fileContents: %s\n", fileContents)

        if err != nil {
            t.Errorf("ReadFile got unexpected error: %v", err)
        }

        if fileContents != testFileContents {
            t.Errorf("ReadFile did not get expected test file contents; got\n%s", fileContents)
        }
    }

b. The function would improve if we used an `io.Reader` or `io.ReadCloser` interface. The result would arguably clean up the code a bit, but more importantly would follow one of Golang's common/predicable patterns. For example:

the function...

    func ReadFile(log *logrus.Logger, f io.ReadCloser, buf []byte) (string, error) {
        defer f.Close()

        var bldr strings.Builder

        for {
            numBytes, err := f.Read(buf)

            bldr.Write(buf[:numBytes])

            log.Debugf("read %d bytes: %s", numBytes, string(buf))

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

the updated unit test...

    func Test(t *testing.T) {
        testFileContents := `1  This royal throne of kings, this sceptred isle,
            2  This earth of majesty, this seat of Mars,
            3  This other Eden, demi-paradise,
            4  This fortress built by Nature for her self
            5  Against infection and the hand of war...`

        testFile, err := os.CreateTemp("", "testFile")
        if err != nil {
            t.Fatal(err)
        }
        defer os.Remove(testFile.Name())

        if _, err := testFile.Write([]byte(testFileContents)); err != nil {
            t.Fatal(err)
        }
        if err := testFile.Close(); err != nil {
            t.Fatal(err)
        }

        log := logrus.New()
        //log.SetLevel(logrus.DebugLevel)

        f, err := os.Open(testFile.Name())
        if err != nil {
            t.Fatal("failed to open test file")
        }

        b := make([]byte, 128)

        fileContents, err := ReadFile(log, f, b)

        fmt.Printf("fileContents: %s\n", fileContents)

        if err != nil {
            t.Errorf("Read got unexpected error: %v", err)
        }

        if fileContents != testFileContents {
            t.Errorf("Read did not get expected test file contents; got\n%s", fileContents)
        }
    }

## 2. Create an original example of "the way to properly wrap types that implement `io.Reader`". (Bodner mentions this at the top of page 236.) Does your example demonstrate how "implementations of `io.Reader` and `io.Writer` are often chained together in a decorator pattern"?
