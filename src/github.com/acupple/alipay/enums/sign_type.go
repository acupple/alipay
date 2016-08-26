package enums

type SignType string

const (
	MD5 SignType = "MD5"

	DSA SignType = "DSA"

	RSA SignType = "RSA"
)

func (p SignType) Humanize() string {
	switch p {
	case MD5:
		return "MD5"
	case DSA:
		return "DSA"
	case RSA:
		return "RSA"
	}
}
