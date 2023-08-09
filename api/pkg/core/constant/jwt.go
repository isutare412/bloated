package constant

type Issuer string

const (
	IssuerNone   Issuer = "none"
	IssuerGoogle Issuer = "google"
)

func (Issuer) Values() []string {
	return []string{
		string(IssuerNone),
		string(IssuerGoogle),
	}
}
