package scanner

import (
	"errors"
	"fmt"
	"os"
	"param-finder/pkg/utils"
	"path/filepath"
	"regexp"
)

type Finding struct {
	path    string
	matches map[string][]string
}

func Scan(dirPath string) ([]Finding, error) {

	var results []Finding

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {

		if err != nil && info.IsDir() {
			return nil
		}

		content, err := utils.ReadFile(path)

		if err != nil {
			return nil
		}

		// AQUI COMEÇARIA A BRINCADEIRA

		result := Finding{
			path:    path,
			matches: FindParameters(content),
		}

		results = append(results, result)

		return nil
	})

	if err != nil {
		return results, errors.New("error while scanning directory " + dirPath)
	}

	return results, nil
}

// "(var|let|const)?\\s[a-zA-Z0-9\\_]{0,40}\\s?=\\s?.*;{0,1}"

// window.location.href
// window.location.search;
// document.URL;
// document.location.search;
// document.location
// window.location
// window.document.location

func FindParameters(sourceCode string) map[string][]string {
	// Define regex patterns
	regexPatterns := map[string]string{
		"Velocity": `request\.getParameter\((?:[^'"])*(?:['"]?([^'"]*)['"]?)\)`,

		"URLSearch":  `new\s*URLSearchParams\([^'"]*\)\.get\(['"]([^'"]*)['"]\)`,
		"WLocationS": `window\.location\.search(?:[^'"])*(?:['"]?([^'"]*)['"]?)`,
		"WLocation":  `window\.location(?:[^'"])*(?:['"]?([^'"]*)['"]?)`,
		"WLocatioH":  `window\.location\.href(?:[^'"])*(?:['"]?([^'"]*)['"]?)`,
		"DLocation":  `(window\.)?document\.location(?:[^'"])*(?:['"]?([^'"]*)['"]?)`,
		"DLocationS": `(window\.)?document\.location\.search(?:[^'"])*(?:['"]?([^'"]*)['"]?)`,
		"DLocationH": `(window\.)?document\.location\.href(?:[^'"])*(?:['"]?([^'"]*)['"]?)`,
		"DURL":       `(window\.)?document\.URL(?:[^'"])*(?:['"]?([^'"]*)['"]?)`,
		"DBaseURI":   `(window\.)?document\.baseURI(?:[^'"])*(?:['"]?([^'"]*)['"]?)`,

		"input":    `<input[^>]+name="([^"]+)"[^>]*>`,
		"select":   `<select[^>]+name="([^"]+)"[^>]*>`,
		"textarea": `<textarea[^>]+name="([^"]+)"[^>]*>`,
		// "gpt3": `(?:var|let|const)?\s+(\w+)\s*(=|:)\s*[^;\n]*new\s*URLSearchParams\([^'"]*(['"][^'"]*['"])`,
	}

	// Initialize a map to store matches for each pattern
	matchesMap := make(map[string][]string)

	// Iterate through regex patterns and find matches
	for key, pattern := range regexPatterns {
		// Compile the regex pattern
		regex := regexp.MustCompile(pattern)

		// Find all matches in the source code
		matches := regex.FindAllStringSubmatch(sourceCode, -1)

		for _, match := range matches {
			if len(match) > 1 && match[1] != "" {
				submatch := match[1]
				// Store matches in the map
				matchesMap[key] = append(matchesMap[key], submatch)
				fmt.Println(submatch)
			}
		}

	}

	return matchesMap
}
