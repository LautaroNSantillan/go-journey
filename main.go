package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMx, hasSPF, sprRecord, hasDMARC, dmarcRecords\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(fmt.Sprintf("Error: could not read from input %v\n", err))
	}
}

func checkDomain(domain string) {
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, errMx := net.LookupMX(domain)
	if errMx != nil {
		log.Printf("Error : %v\n", errMx)
	}

	if len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, errTxt := net.LookupTXT(domain)
	if errTxt != nil {
		log.Printf("Error : %v\n", errTxt)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, errDmarc := net.LookupTXT("_dmarc." + domain)
	if errDmarc == nil {
		log.Printf("Error : %v\n", errDmarc)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v \n", domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
