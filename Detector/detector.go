package Detector

import (
	"bufio"
	"errors"
	"os"

	"github.com/irfansharif/cfilter"
)

//Detector represents a prohibited detector
type Detector struct {
	cf *cfilter.CFilter //CFilter used for search words from a probabilistic data store
}

// New returns a new Detector object. And it will bind a CFilter when created
func New() *Detector {
	dt := new(Detector)
	dt.cf = cfilter.New()
	return dt
}

//Init reads the words file and store all elements into the CFilter.
//filepath is your words file path
//returns error when the words file not exists or scan the file occur error
func (d Detector) Init(filepath string) error {
	if filepath == "" {
		return errors.New("Words file not specified")
	}
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		d.cf.Insert([]byte(scanner.Text()))
	}

	if err := scanner.Err();err != nil {
		return err
	}
	return nil
}

// Lookup checks if an element (in byte-array form) exists in the Cuckoo
// Filter, returns true if found and false otherwise.
func (d Detector) Lookup(words string) bool {
	return d.cf.Lookup([]byte(words))
}
