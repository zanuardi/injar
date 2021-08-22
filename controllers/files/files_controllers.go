package files

import (
	"injar/usecase/files"
	"net/http"

	controller "injar/controllers"

	echo "github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

type FileController struct {
	FileUC files.Usecase
	DB     *gorm.DB
}

func NewFileController(uc files.Usecase, db *gorm.DB) *FileController {
	return &FileController{
		DB: db,
		// FilesUC: uc,
	}
}

func (ctrl *FileController) Store(c echo.Context) error {
	// ctx := c.Request().Context()
	fileType := c.FormValue("file_type")
	file, err := c.FormFile("file")
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	// fileUC := usecase.NewFileUC(ctx, DB)
	res, err := ctrl.FileUC.Store(fileType, file)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, res)
}
