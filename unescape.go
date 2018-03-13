package httptreemux

import "net/url"

func unescape(path string) (string, error) {
	return url.PathUnescape(path)
}
