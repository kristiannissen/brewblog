package pkg

/*
 * Shared interface methods
 */
type Service interface {
	Get(url string) ([]byte, error)
	List() ([]Blob, error)
}
