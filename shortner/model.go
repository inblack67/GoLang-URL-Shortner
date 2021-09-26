package shortner

type MRedirect struct {
	Code      string `json:"code"`
	URL       string `json:"url" validate:"empty=false & format=url"`
	CreatedAt string `json:"createdAt"`
}
