package member

// CreateMemberRequest represents the request body for creating a member.
type CreateMemberRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

// GetMemberRequest represents the request parameters for getting a member.
// ID is extracted from the URL path parameter.
type GetMemberRequest struct {
	ID int `param:"id" validate:"required,min=1"`
}
