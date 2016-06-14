package request

// BaseAPIResponse is common api response data
type BaseAPIResponse struct {
	Code    int    `json:"code"`
	Version string `json:"version"`
	Status  string `json:"status"`
}
