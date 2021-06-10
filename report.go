package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"testing"
)

var report [][]string

func initReportWriter() func() {
	header := []string{"ID", "Status", "Related Assertions from Spec", "Comments"}

	writer := csv.NewWriter(os.Stdout)

	err := writer.Write(header)
	if err != nil {
		fmt.Printf("Error writing header to report file: %s", err)
		os.Exit(1)
	}

	return func() {
		fmt.Println("==================== REPORT ====================")
		for _, record := range report {
			err := writer.Write(record)
			if err != nil {
				fmt.Printf("Error writing the report: %s", err)
				os.Exit(1)
			}
		}

		writer.Flush()
		fmt.Println("================================================")
		err = writer.Error()
		if err != nil {
			fmt.Printf("Error flushing the report file: %s", err)
			os.Exit(1)
		}
	}

}

func writeReportLine(id, status, assertions, comments string) {
	report = append(report, []string{id, status, assertions, comments})
}

func writeTestResult(id, assertions, comments string, t *testing.T) {
	if strings.Contains(id, ",") {
		t.Fatalf("ID cell should not contain commas.")
	}
	if strings.Contains(assertions, ",") {
		t.Fatalf("Assertions cell should not contain commas. Use space as separator.")
	}
	if t.Failed() {
		writeReportLine(id, "failed", assertions, comments)
	} else if t.Skipped() {
		writeReportLine(id, "skipped", assertions, comments)
	} else {
		writeReportLine(id, "passed", assertions, comments)
	}
}