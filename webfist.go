// Package webfist implements WebFist.
package webfist

// Storage is the interface implemented by backends.
type Storage interface {
	PutEmail(*EmailAddr, *Email) error
	Emails(*EmailAddr) ([]*Email, error)

	// StatEncryptedBlob returns the size of the encrypted blob on
	// disk. sha1 is lowercase hex. The err will be os.ErrNotExist
	// if the file is doesn't exist.
	StatEncryptedEmail(sha1 string) (size int, err error)

	// EncryptedEmail returns the encrypted email with
	// the the encrypted digest of sha1 (lowercase hex).
	// The err will be os.ErrNotExist if the file is doesn't
	// exist.
	EncryptedEmail(sha1 string) ([]byte, error)
}

// Payload from: http://tools.ietf.org/html/draft-ietf-appsawg-webfinger
// TODO: Make this type pretty and more native, not just a bag of properties.
type WebFingerResponse struct {
  JSON map[string]interface{}
}

// Lookup performs a WebFinger query for an email address and returns all known
// data for that address. Implementations may do standard WebFinger lookups over
// the network, fallback to using the WebFist network, or use local storage to
// map email address to WebFinger response.
type Lookup interface {
  WebFinger(emailAddr string) (*WebFingerResponse, error)
}
