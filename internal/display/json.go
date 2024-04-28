package display

import (
	"encoding/json"
	"fmt"
	"hurryup/internal/speed"
)

func JSON(tester speed.Tester) error {
	data := speed.Data{ServerData: tester.GetServerData(), TesterData: tester.GetTesterData()}
	output, err := json.Marshal(data)

	if err != nil {
		return fmt.Errorf("could not serialize json: %w", err)
	}

	fmt.Println(string(output))

	return nil
}
