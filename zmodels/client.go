package zmodels

// Context: host url to be used on the request made using the client
// Headers: headers to be included in every call made using the client.
// The headers added on ZRequest have priority over the client ones.
type ZClient struct {
	ContextUrl string
	Headers    map[string]string
}
