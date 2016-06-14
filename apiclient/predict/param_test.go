package predict

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		projectID string
		modelID   string
		data      interface{}
		ok        bool
		errString string
	}{
		{"project", "model", "data", true, ""},
		{"", "model", "data", false, "ProjectID"},
		{"project", "", "data", false, "ModelID"},
		{"project", "model", nil, false, "Data"},
		{"", "", nil, false, "ProjectID"},
		{"project", "", nil, false, "ModelID"},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("%+v", tt)

		p := Param{
			ProjectID: tt.projectID,
			ModelID:   tt.modelID,
			Data:      tt.data,
		}

		ok, name := p.IsValid()
		assert.Equal(tt.ok, ok, target)
		if !tt.ok {
			assert.Equal(tt.errString, name, target)
		}
	}

}
