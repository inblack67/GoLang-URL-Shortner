package shortner

import (
	"encoding/json"
	"io"
)

func (redirects TRedirects) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(redirects)
}

func (redirect *MRedirect) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(redirect)
}
