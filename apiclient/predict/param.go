package predict

// Param contains api request data
type Param struct {
	ProjectID string
	ModelID   string
	Data      interface{}
}

// IsValid checks parameters
func (p Param) IsValid() (bool, string) {
	switch {
	case p.ProjectID == "":
		return false, "ProjectID"
	case p.ModelID == "":
		return false, "ModelID"
	case p.Data == nil:
		return false, "Data"
	}

	return true, ""
}
