package scanner

type ParsedResult struct {
	Path   string
	Params []string
}

func isInArray(str string, array []string) bool {
	for _, value := range array {
		if value == str {
			return true
		}
	}
	return false
}

func Parse(results []Finding) []ParsedResult {

	var newResults []ParsedResult

	for _, file := range results {
		result := ParsedResult{
			Path: file.path,
		}

		for _, match := range file.matches {
			for _, param := range match {
				// Filtering repeated items
				if !(isInArray(param, result.Params)) && param != "&" {
					result.Params = append(result.Params, param)
				}
			}
		}

		newResults = append(newResults, result)

	}

	return newResults

}
