package shortener

type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Save(redirect *Redirect) error
}
