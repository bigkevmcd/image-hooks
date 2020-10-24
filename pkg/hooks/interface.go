package hooks

// PushEvent values return the image that is to be inserted into the file.
type PushEvent interface {
	PushedImageURL() string
	EventRepository() string
	EventTag() string
}

// PushEventParser parses the specifics of a hook request into a body.
type PushEventParser func(payload []byte) (PushEvent, error)
