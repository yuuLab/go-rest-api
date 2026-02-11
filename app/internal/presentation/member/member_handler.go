package member

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/yuuLab/go-rest-api/internal/application/usecase"
	"github.com/yuuLab/go-rest-api/internal/presentation/httperror"
)

// Handler handles member-related HTTP requests.
type Handler struct {
	registerUseCase usecase.RegisterMemberUseCaseI
	getUseCase      usecase.GetMemberUseCaseI
}

// NewHandler creates a new member handler.
func NewHandler(
	registerUseCase usecase.RegisterMemberUseCaseI,
	getUseCase usecase.GetMemberUseCaseI,
) *Handler {
	return &Handler{
		registerUseCase: registerUseCase,
		getUseCase:      getUseCase,
	}
}

// Create handles POST /members - creates a new member.
func (h *Handler) Create(c echo.Context) error {
	var req CreateMemberRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httperror.ErrorResponse{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	// Validate request
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httperror.ErrorResponse{
			Message: "Validation failed",
			Error:   err.Error(),
		})
	}

	// Execute use case
	input := usecase.RegisterMemberInput{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	output, err := h.registerUseCase.Execute(c.Request().Context(), input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httperror.ErrorResponse{
			Message: "Failed to create member",
			Error:   err.Error(),
		})
	}

	// Convert to response
	response := MemberResponse{
		ID:        output.ID,
		FirstName: output.FirstName,
		LastName:  output.LastName,
		CreatedAt: output.CreatedAt,
	}

	return c.JSON(http.StatusCreated, response)
}

// Get handles GET /members/:id - retrieves a member by ID.
func (h *Handler) Get(c echo.Context) error {
	// Get ID from path parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httperror.ErrorResponse{
			Message: "Invalid member ID",
			Error:   "ID must be a number",
		})
	}

	// Execute use case
	input := usecase.GetMemberInput{
		ID: id,
	}

	output, err := h.getUseCase.Execute(c.Request().Context(), input)
	if err != nil {
		// Check if member not found
		if err.Error() == "member not found" {
			return c.JSON(http.StatusNotFound, httperror.ErrorResponse{
				Message: "Member not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, httperror.ErrorResponse{
			Message: "Failed to get member",
			Error:   err.Error(),
		})
	}

	// Convert to response
	response := MemberResponse{
		ID:        output.ID,
		FirstName: output.FirstName,
		LastName:  output.LastName,
		CreatedAt: output.CreatedAt,
	}

	return c.JSON(http.StatusOK, response)
}
