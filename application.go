package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/armor"
	"github.com/ProtonMail/go-crypto/openpgp/packet"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var shouldExportKey bool
var exportFileName string

func init() {
	flag.BoolVar(&shouldExportKey, "save", false, "Save the public key to disk")
	flag.StringVar(&exportFileName, "filename", "public_key.pgp", "The filename the public key should be saved as")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Protondate 0.0.1 by Doctor Chaos\n\nUsage: %s <email address>\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("Error: No email address has been provided. Please provide a email address!")
	}

	emailAddress := flag.Arg(0)

	url := fmt.Sprintf("https://api.protonmail.ch/pks/lookup?op=get&search=%s", emailAddress)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if string(body) == "No key found" {
		log.Fatalf("No PGP key found for \"%s\" the e-mail address might not be used on protonmail\n", emailAddress)
		return
	}

	if shouldExportKey {
		if err != nil {
			log.Printf("Error: Unable to save response! (%s)\n", err)
		}
		err = ioutil.WriteFile(exportFileName, body, 0777)
		if err != nil {
			log.Printf("Error: Unable to save response! (%s)\n", err)
		}
	}

	block, err := armor.Decode(bytes.NewReader(body))

	if err != nil {
		log.Fatal(err)
		return
	}

	if block.Type != openpgp.PublicKeyType {
		log.Fatalf("The PGP key belonging to \"%s\" is not a public key! WTF?\n", emailAddress)
		return
	}

	entity, err := openpgp.ReadEntity(packet.NewReader(block.Body))
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("The protonmail account belonging to \"%s\" was potentially created on %s\n", flag.Arg(0), entity.PrimaryKey.CreationTime)
}
