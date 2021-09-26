package shortner

type SRedirect interface {
	Find(code string) (MRedirect, error)
	Store(url string) (bool, error)
}
