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
