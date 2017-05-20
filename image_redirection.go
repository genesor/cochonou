package cochonou

// ImageRedirection is the struct that represents a redirection for an image.
type ImageRedirection struct {
	ID        string
	SubDomain string
	URL       string
}

// ImageRedirectionStore is the interface used to store ImageRedirections
type ImageRedirectionStore interface {
	Save(redir *ImageRedirection) error
	All() ([]ImageRedirection, error)
}
