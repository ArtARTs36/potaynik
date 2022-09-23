package viewer

type ViewParams struct {
	SecretKey   string
	AuthFactors map[string]string
	UserIP      string
}
