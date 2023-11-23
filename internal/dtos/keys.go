package dtos

type ContextKey string

const (
	ContextKeyAccountID         ContextKey = "accountID"
	ContextKeyIpAddress         ContextKey = "ipAddress"
	ContextKeyLang              ContextKey = "lang"
	ContextKeyLeadingWhitespace ContextKey = "leadingWhitespace"
	ContextKeyRepeatRecipients  ContextKey = "repeatRecipientsKey"
	ContextKeyRequestID         ContextKey = "requestId"
	ContextKeyTokenInfo         ContextKey = "tokenInfo"
	ContextKeyUserAgent         ContextKey = "userAgent"
)
