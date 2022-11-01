package zmodels

// Request data carrying struct for MakeRequest function.
// If `IsBytesBody` is set to true, MakeRequest function will expect
// for the request body to be of type []byte. This is to avoid marshalling
// and buffering the same bytes array twice.
type ZRequest[T any] struct {
	Url               string
	Headers           map[string]string
	Body              T
	UrlParams         map[string]string
	QueryParams       map[string][]string
	IsBytesBody       bool
	IgnoreJsonParsing bool
}
