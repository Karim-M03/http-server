package response


import "encoding/json"

// used server side to facilitate the creation of response
type TmpResponse struct {
    Status  int `json:"status"`
    Message string `json:"message"`
    Data    any    `json:"data,omitempty"`
}

func (r TmpResponse) ToJSON() ([]byte, error) {
    return json.Marshal(r)
}