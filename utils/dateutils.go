package utils

import "time"

func TransformDate(input string) (string, error) {
	inputFormat := "2/1/2006"
	outputFormat := "02-01-2006"

	parsedTime, err := time.Parse(inputFormat, input)
	if err != nil {
		return "", err
	}

	output := parsedTime.Format(outputFormat)
	return output, nil
}
