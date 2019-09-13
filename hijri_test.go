package hijri_test

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/RadhiFadlillah/go-hijri"
)

var hijriTests []comparisonTest

type comparisonTest struct {
	Gregorian string
	Hijri     string
}

func init() {
	// Open test file
	f, err := os.Open("test/hijri.csv")
	if err != nil {
		log.Fatalf("failed to open Hijri test file: %v\n", err)
	}
	defer f.Close()

	// Parse test file
	csvReader := csv.NewReader(f)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed to parse Hijri test file: %v\n", err)
		}

		hijriTests = append(hijriTests, comparisonTest{
			Gregorian: record[0],
			Hijri:     record[1],
		})
	}
}

func TestToHijri(t *testing.T) {
	if len(hijriTests) == 0 {
		t.Fatal("no tests available for Hijri")
	}

	for _, item := range hijriTests {
		date, _ := time.Parse("2006-01-02", item.Gregorian)
		year, month, day := hijri.ToHijri(date)
		result := fmt.Sprintf("%04d-%02d-%02d", year, month, day)

		if result != item.Hijri {
			t.Errorf("%s: want %s got %s\n", item.Gregorian, item.Hijri, result)
		}
	}
}

func TestFromHijri(t *testing.T) {
	if len(hijriTests) == 0 {
		t.Fatal("no tests available for Hijri")
	}

	for _, item := range hijriTests {
		var hYear, hMonth, hDay int
		fmt.Sscanf(item.Hijri, "%d-%d-%d", &hYear, &hMonth, &hDay)

		date := hijri.FromHijri(hYear, hMonth, hDay)
		result := date.Format("2006-01-02")

		if result != item.Gregorian {
			t.Errorf("%s: want %s got %s\n", item.Hijri, item.Gregorian, result)
		}
	}
}
