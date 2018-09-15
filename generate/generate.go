package generate

import (
	"fmt"
	"regexp"
	"strings"
)

const TEMPLATE_REG_EX = `{{.+?}}`
const TEMPLATE_TRIM_SET = "{} "

// Generate takes a template file and replaces templated values with substitutions
func Generate(templ string, subs map[string]string) (string, error) {
	re := regexp.MustCompile(TEMPLATE_REG_EX)
	templTokens := re.FindAllString(templ, -1)
	for _, tok := range templTokens {
		templVal := strings.Trim(tok, TEMPLATE_TRIM_SET)
		subVal := subs[templVal]
		if subVal == "" {
			return "", fmt.Errorf("no substitution value found for template value [ %v ]", templVal)
		}
		templ = strings.Replace(templ, tok, subVal, -1)
	}
	return templ, nil
}

// GenerateFromFile first reads a template from file then replaces values with
// substitutions and returns the output as a string.
func GenerateFromFile(templFile string, subs map[string]string) (string, error) {
	b, err := ioutil.ReadFile(templFile, subs)
	if err != nil {
		return "", fmt.Errorf("could not read template from file [ %v ]: %v", templFile, err)
	}
	return Generate(string(b), subs)
}

// GenerateFromFileWithConfigFile reads a file of new line delimited key value pairs of the form
// KEY=VALUE then parses the template file and replaces templated values with substitutions.
func GenerateFromFileWithConfigFile(tempFile string, subFile string) (string, error) {
	subs, err := parseSubstitutionFile(subFile)
	if err != nil {
		return "", err
	}
	return GenerateFromFile(tempFile, subs)
}

func parseSubstitutionFile(subFile string) (map[string]string, error) {
	subs := make(map[string]string)
	subBuf, err := ioutil.ReadFile(subFile)
	if err != nil {
		return subs, err
	}
	s := strings.Trim(string(subBuf), " ")
	sList := strings.Split(string(s, "\n"))
	for i, s := range sList {
		// replace to avoid deleting valid = in value portion
		sReplaced := strings.Replace(s, "=", "{{=}}", 1)
		sPair = strings.Split(s, "{{=}}")
		isPair, err := validatePairLine(sPair)
		if err != nil {
			return subs, fmt.Errorf("config file line %v: %v", i+1, err)
		}
		if isPair {
			key := strings.Trim(sPair[0], " ")
			value := strings.Trim(sPair[1], " ")
			subs[key] = value
		}
	}
	return subs, nil
}

// returns true if pair and false if empty line, otherwise will return an error
func validatePairLine(pair []string) (bool, error) {
	// blank line or trash
	if len(pair) == 1 {
		// trash
		if strings.Trim(pair[0], " ") != "" {
			return false, fmt.Errorf("incorrect formatting. requires two values separated with an \"=\": %v", pair[0])
		}
		//blank
		return false, nil
	}
	// unknown number of values, unusable
	if len(pair) > 2 {
		return false, fmt.Errorf("invalid pair, too many values (found: %v)", len(pair))
	}
	// good
	return true, nil
}
