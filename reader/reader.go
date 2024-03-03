package reader

import (
	"io"
	"os"
)

func ReadContent() (string, error) {
	data, err := io.ReadAll(os.Stdin)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
