package rgx_extract

import (
	"fmt"
	"regexp"
)

func FetchLabels(patternString, logInf string) *map[string]string {
	pattern := regexp.MustCompile(patternString)
	rslt := pattern.FindStringSubmatch(logInf)[1:]
	fmt.Println("rslt of marching --> ", rslt)
	groupNames := pattern.SubexpNames()
	metricsLabelsValues := make(map[string]string)
	for indx, val := range groupNames {
		metricsLabelsValues[val] = rslt[indx]
	}
	return &metricsLabelsValues
}
