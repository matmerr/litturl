package server

// wow generics would be neat for this
type database interface {
	Get(string) (URLTranslation, error)
	Put(string, URLTranslation) error
	NewUser(string, string) error
	IsUser(user) (bool, error)
}
