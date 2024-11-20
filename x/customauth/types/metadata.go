package types

import (
	"encoding/json"
	"strings"

	"github.com/gogo/protobuf/jsonpb"
)

func DecodeMetadata(memo string) (*CustomAuthMetadata, error) {
	if len(memo) == 0 {
		return nil, nil
	}
	d := make(map[string]interface{})
	err := json.Unmarshal([]byte(memo), &d)
	if err != nil {
		return nil, nil
	}

	if d["custom_auth"] == nil {
		return nil, nil
	}

	m := &CustomAuthMetadata{}
	err = jsonpb.Unmarshal(strings.NewReader(memo), m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
