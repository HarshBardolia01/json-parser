package filereader

import "os"

func ReadContent(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(data), nil
}