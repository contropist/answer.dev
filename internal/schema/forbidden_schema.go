package schema

const (
	ForbiddenReasonTypeInactive      = "inactive"
	ForbiddenReasonTypeUrlExpired    = "url_expired"
	ForbiddenReasonTypeUserSuspended = "suspended"
)

// ForbiddenResp forbidden response
type ForbiddenResp struct {
	// forbidden reason type
	Type string `json:"type" enums:"inactive,url_expired"`
}
