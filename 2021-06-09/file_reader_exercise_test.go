package localFileReader

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

	// the rest of the unit test...
}
