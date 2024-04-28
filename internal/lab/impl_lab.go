package lab

import (
	"fmt"
	"net/http"

	"github.com/Magum23/wac-lab-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Copy following section to separate file, uncomment, and implement accordingly
// DeleteLabAnalysis - Delete laboratory analysis by ID
func (i *implLabAPI) DeleteLabAnalysis(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[LaboratoryAnalysis])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	analysisId := ctx.Param("analysisId")
	err := db.DeleteDocument(ctx, analysisId)

	switch err {
	case nil:
		ctx.AbortWithStatus(http.StatusNoContent)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Analysis not found",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to delete analysis from database",
				"error":   err.Error(),
			})
	}
}

// DeleteSampleEvidence - Delete sample evidence by ID
func (i *implLabAPI) DeleteSampleEvidence(ctx *gin.Context) {

	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[SampleEvidence])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	evidenceId := ctx.Param("evidenceId")
	err := db.DeleteDocument(ctx, evidenceId)

	switch err {
	case nil:
		ctx.AbortWithStatus(http.StatusNoContent)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Evidence not found",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to delete evidence from database",
				"error":   err.Error(),
			})
	}

}

// GetLabAnalysis - Get all laboratory analyses
func (i *implLabAPI) GetLabAnalysis(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// GetLabAnalysisById - Get laboratory analysis by ID
func (i *implLabAPI) GetLabAnalysisById(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// GetSampleEvidence - Get all sample evidences
func (i *implLabAPI) GetSampleEvidence(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// GetSampleEvidenceById - Get sample evidence by ID
func (i *implLabAPI) GetSampleEvidenceById(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// LabAnalysis - Create a new laboratory analysis
func (i *implLabAPI) CreateLabAnalysis(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[LaboratoryAnalysis])

	fmt.Println("DB: ", db, ok, value.(db_service.DbService[LaboratoryAnalysis]))

	fmt.Println("DB: ", db, ok, value)
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	analysis := LaboratoryAnalysis{}
	err := ctx.BindJSON(&analysis)

	fmt.Println("Error: ", err)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	if analysis.Id == "" {
		analysis.Id = uuid.New().String()
	}

	err = db.CreateDocument(ctx, analysis.Id, &analysis)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusCreated,
			analysis,
		)
	case db_service.ErrConflict:
		ctx.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Analysis already exists",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create ambulance in database",
				"error":   err.Error(),
			},
		)
	}
}

// SampleEvidence - Create a new sample evidence
func (i *implLabAPI) CreateSampleEvidence(ctx *gin.Context) {

	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[SampleEvidence])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	evidence := SampleEvidence{}
	err := ctx.BindJSON(&evidence)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	if evidence.Id == "" {
		evidence.Id = uuid.New().String()
	}

	err = db.CreateDocument(ctx, evidence.Id, &evidence)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusCreated,
			evidence,
		)
	case db_service.ErrConflict:
		ctx.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Evidence already exists",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create evidence in database",
				"error":   err.Error(),
			},
		)
	}

}

// UpdateLabAnalysis - Update laboratory analysis by ID
func (i *implLabAPI) UpdateLabAnalysis(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// UpdateSampleEvidence - Update sample evidence by ID
func (i *implLabAPI) UpdateSampleEvidence(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}
