package httptreemux

import (
	"context"
)

// ContextParams returns the params map associated with the given context if one exists. Otherwise, an empty map is returned.
func ContextParams(ctx context.Context) map[string]string {
	if p, ok := ctx.Value(paramsContextKey).(map[string]string); ok {
		return p
	}
	return map[string]string{}
}

// AddParamsToContext inserts a parameters map into a context using
// the package's internal context key. Clients of this package should
// really only use this for unit tests.
func AddParamsToContext(ctx context.Context, params map[string]string) context.Context {
	return context.WithValue(ctx, paramsContextKey, params)
}

type contextKey int

// paramsContextKey is used to retrieve a path's params map from a request's context.
const paramsContextKey contextKey = 0
