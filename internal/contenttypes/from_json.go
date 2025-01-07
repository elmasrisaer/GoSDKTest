package contenttypes

import (
	"fmt"
	"github.com/elmasrisaer/GoSDKTest/internal/unmarshal"
)

func FromJson(data []byte, target any) error {
	err := unmarshal.Unmarshal(data, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body into struct: %v", err)
	}
	return nil
}
