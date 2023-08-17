package ucontexts

import "context"

const CtxKeyLogId = "U_LOG_ID"

// SetLogId set logId to context
func SetLogId(ctx context.Context, logId string) context.Context {
	return context.WithValue(ctx, CtxKeyLogId, logId)
}

// GetLogId get logId from context, return (logId, ok)
func GetLogId(ctx context.Context) (string, bool) {
	if ctx == nil {
		return "", false
	}

	v := ctx.Value(CtxKeyLogId)
	if v == nil {
		return "", false
	}

	switch t := v.(type) {
	case string:
		return t, true
	case *string:
		if t == nil {
			return "", false
		}
		return *t, true
	}
	return "", false
}

func OptLogId(ctx context.Context) string {
	v, _ := GetLogId(ctx)
	if v == "" {
		return "-"
	}
	return v
}
