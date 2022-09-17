package internal

func GenerateRequestUrl(context, uri string) string {
	contextLastIndex := len(context) - 1

	if string(context[contextLastIndex]) == "/" {
		context = context[:contextLastIndex]
	}

	if string(uri[0]) != "/" {
		uri = "/" + uri
	}

	return context + uri
}
