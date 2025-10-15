package ast

import (
	"encoding/json"
	"fmt"
)

func (member Member) PropertyType() (PropertyType, error) {
	var result PropertyType
	if len(member.TypePayload) == 0 {
		return result, fmt.Errorf("member %s does not contain property type information", member.Name)
	}
	unmarshalErr := json.Unmarshal(member.TypePayload, &result)
	if unmarshalErr != nil {
		return result, fmt.Errorf("unable to decode property type for %s: %w", member.Name, unmarshalErr)
	}
	return result, nil
}

func (member Member) ConstType() (Identifier, error) {
	var result Identifier
	if len(member.TypePayload) == 0 {
		return result, fmt.Errorf("member %s does not contain const type information", member.Name)
	}
	unmarshalErr := json.Unmarshal(member.TypePayload, &result)
	if unmarshalErr != nil {
		return result, fmt.Errorf("unable to decode const type for %s: %w", member.Name, unmarshalErr)
	}
	return result, nil
}
