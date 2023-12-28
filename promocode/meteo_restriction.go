package promocode

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type MeteoRestriction struct {
	Is   string
	Temp struct {
		Gt int
	}
}

func (r MeteoRestriction) Validate(arg Arguments) (bool, error) {
	if r.Is != "" && r.Is != arg.MeteoStatus {
		return false, fmt.Errorf("invalid meteo status: expected %v (got %v)", r.Is, arg.MeteoStatus)
	}

	if arg.MeteoTemp < r.Temp.Gt {
		return false, fmt.Errorf("invalid temperature: should be greater than %v (got %v)", r.Temp.Gt, arg.Age)
	}

	return true, nil
}

func (d *MeteoRestriction) UnmarshalJSON(data []byte) error {
	var result map[string]any
	err := json.Unmarshal(data, &result)
	if err != nil {
		return fmt.Errorf("failed to parse json: %v", err)
	}

	d.Is = result["is"].(string)
	temp := result["temp"].(map[string]any)
	gtStr := temp["gt"].(string)
	d.Temp.Gt, err = strconv.Atoi(gtStr)

	if err != nil {
		return fmt.Errorf("failed to parse gt: %v", err)
	}

	return err
}
