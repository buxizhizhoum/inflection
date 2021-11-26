package inflection

import (
	"regexp"
	"strings"
)

func Underscore(s string) string {
	for _, reStr := range []string{`([A-Z]+)([A-Z][a-z])`, `([a-z\d])([A-Z])`} {
		re := regexp.MustCompile(reStr)
		s = re.ReplaceAllString(s, "${1}_${2}")
	}
	return strings.ToLower(s)
}

func Camelize(s string, uppercaseFirstLetter bool) string {
	replFunc := func(w string) string {
		if strings.HasPrefix(w, "_") && !strings.HasPrefix(w, "__"){
			return strings.ToUpper(w[1:])
		}
		return strings.ToUpper(w)
	}

	if uppercaseFirstLetter {
		re := regexp.MustCompile(`(?:^|_)(.)`)
		return re.ReplaceAllStringFunc(s, replFunc)
	} else {
		return strings.ToLower(string(s[0])) + Camelize(s, true)[1:]
	}
}
