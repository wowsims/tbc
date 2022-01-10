package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	outDir := flag.String("outDir", "", "Path to output directory for writing generated .go files.")
	flag.Parse()

	if *outDir == "" {
		panic("outDir flag is required!")
	}

	tooltipsDB := getWowheadTooltipsDB()

	gemDeclarations := getGemDeclarations()
	gemsData := make([]GemData, len(gemDeclarations))
	for idx, gemDeclaration := range gemDeclarations {
		gemData := GemData{
			Declaration: gemDeclaration,
			Response:    getWowheadItemResponse(gemDeclaration.ID, tooltipsDB),
		}
		//log.Printf("\n\n%+v\n", gemData.Response)
		gemsData[idx] = gemData
	}
	sort.SliceStable(gemsData, func(i, j int) bool {
		return gemsData[i].Response.Name < gemsData[j].Response.Name
	})
	writeGemFile(*outDir, gemsData)

	itemResponses := make([]WowheadItemResponse, len(ItemDeclarations))
	for idx, itemDeclaration := range ItemDeclarations {
		itemResponse := getWowheadItemResponse(itemDeclaration.ID, tooltipsDB)
		//fmt.Printf("\n\n%+v\n", itemResponse)
		itemResponses[idx] = itemResponse
	}
	writeItemFile(*outDir, ItemDeclarations, itemResponses)
}

func getGemDeclarations() []GemDeclaration {
	gemsData := readCsvFile("./assets/item_data/all_gem_ids.csv")

	// Ignore first line
	gemsData = gemsData[1:]

	gemDeclarations := make([]GemDeclaration, len(gemsData))
	for i, gemsDataRow := range gemsData {
		gemID, err := strconv.Atoi(gemsDataRow[0])
		if err != nil {
			log.Fatal("Invalid gem ID: " + gemsDataRow[0])
		}
		declaration := GemDeclaration{
			ID: gemID,
		}

		for _, override := range GemDeclarationOverrides {
			if override.ID == gemID {
				declaration = override
				break
			}
		}

		gemDeclarations[i] = declaration
	}

	// Add any declarations that were missing from the csv file.
	for _, overrideGemDeclaration := range GemDeclarationOverrides {
		found := false
		for _, gemDecl := range gemDeclarations {
			if gemDecl.ID == overrideGemDeclaration.ID {
				found = true
				break
			}
		}
		if !found {
			gemDeclarations = append(gemDeclarations, overrideGemDeclaration)
		}
	}

	return gemDeclarations
}

// Returns the prefetched list of all wowhead tooltips.
// Maps item IDs to tooltip strings.
func getWowheadTooltipsDB() map[int]string {
	file, err := os.Open("./assets/item_data/all_item_tooltips.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	db := make(map[int]string)
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
		if i == 1 {
			// Ignore first line
			continue
		}

		line := scanner.Text()

		itemIDStr := line[:strings.Index(line, ",")]
		itemID, err := strconv.Atoi(itemIDStr)
		if err != nil {
			log.Fatal("Invalid item ID: " + itemIDStr)
		}

		tooltip := line[strings.Index(line, "{"):]
		db[itemID] = tooltip
	}

	return db
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
