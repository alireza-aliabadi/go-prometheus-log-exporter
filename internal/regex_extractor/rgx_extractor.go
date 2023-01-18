package rgx_extract

import (
	"regexp"
)

func FetchLabels(patternString, logInf string) map[string]string {
	pattern := regexp.MustCompile(patternString)
	rslt := pattern.FindStringSubmatch(logInf)[1:]
	groupNames := pattern.SubexpNames()
	metricsLabelsValues := map[string]string{}
	for indx, val := range groupNames[1:] {
		metricsLabelsValues[val] = rslt[indx]
	}
	return metricsLabelsValues
}
