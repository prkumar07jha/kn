package source

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"

	"kn/se/pkg/model"
)

// Parse parses source files and return list of source links.
func Parse() []model.SourceLink {
	log.Println("Parser started")

	file, err := os.Open("assets/initial_sources.json")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Error while opening file"))
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Error while reading file"))
	}

	var sourceLinks []model.SourceLink
	err = json.Unmarshal(bytes, &sourceLinks)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Error while parsing JSON"))
	}

	return sourceLinks
}