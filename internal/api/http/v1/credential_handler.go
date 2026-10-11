package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/vucongthanh92/courier-user-svc/internal/domain/interfaces"
	"github.com/vucongthanh92/courier-user-svc/internal/domain/models"

	utilsCommon "github.com/vucongthanh92/go-base-utils/helper/http_common"
	utilsError "github.com/vucongthanh92/go-base-utils/helper/http_error"
	baseUtils "github.com/vucongthanh92/go-base-utils/helper/utils"
)

type CredentialHandler struct {
	credentialService interfaces.AuthCredentialServiceI
}

func InitCredentialHandler(credentialService interfaces.AuthCredentialServiceI) *CredentialHandler {
	return &CredentialHandler{
		credentialService: credentialService,
	}
}

// API GeneratePassword godoc
// @Tags Credential
// @Summary Generate password for user
// @Accept json
// @Produce json
// @Param params body models.GeneratePasswordRequest true "GeneratePasswordRequest"
// @Router /api/v1/user/pwd/generate [post]
// @Success	200 {object} utilsCommon.SuccessResponse[string]
func (h *CredentialHandler) GeneratePassword(c *gin.Context) {

	// step 1. bind request body to struct and validate
	req := models.GeneratePasswordRequest{}
	if err := utilsCommon.GetBodyParamsHTTP(c, &req); err != nil {
		return
	}

	// Validate request parameters
	if err := utilsCommon.ValidatorParams(req); err != nil {
		utilsError.InitErrorBuilder(c).
			SetStatus(http.StatusBadRequest).
			SetArrayError(err).
			ExposeHttpError(c)
		return
	}

	// step 2. get userID from token and set to request struct
	claims := c.Value("authClaims").(jwt.MapClaims)
	req.UserID = baseUtils.ParseUserID(claims["sub"])

	// step 3. call use case to generate password
	ctx := baseUtils.SetHeaderByKey(c, "headers")
	if resErr := h.credentialService.SetPassword(ctx, req); resErr != nil {
		resErr.ExposeHttpError(c)
		return
	}

	c.JSON(http.StatusOK, utilsCommon.NewSuccessResponse("Password generated successfully"))
}
