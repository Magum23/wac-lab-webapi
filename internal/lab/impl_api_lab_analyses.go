package lab

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Copy following section to separate file, uncomment, and implement accordingly
// CreateLabAnalysis - Create a new laboratory analysis
func (t *implLabAnalysesAPI) CreateLabAnalysis(ctx *gin.Context) {
	updateLaboratoryFunc(ctx, func(c *gin.Context, laboratory *Laboratory) (*Laboratory, interface{}, int) {
		var laboratoryAnalysis LaboratoryAnalysis

		if err := c.ShouldBindJSON(&laboratoryAnalysis); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		if laboratoryAnalysis.Id == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Analysis ID is required",
			}, http.StatusBadRequest
		}

		if laboratoryAnalysis.Id == "" || laboratoryAnalysis.Id == "@new" {
			laboratoryAnalysis.Id = uuid.NewString()
		}

		conflictIndx := slices.IndexFunc(laboratory.LabAnalyses, func(la LaboratoryAnalysis) bool {
			return laboratoryAnalysis.Id == la.Id
		})

		if conflictIndx >= 0 {
			return nil, gin.H{
				"status":  http.StatusConflict,
				"message": "Analysis already exists",
			}, http.StatusConflict
		}

		laboratory.LabAnalyses = append(laboratory.LabAnalyses, laboratoryAnalysis)
		analysisIndx := slices.IndexFunc(laboratory.LabAnalyses, func(la LaboratoryAnalysis) bool {
			return laboratoryAnalysis.Id == la.Id
		})
		if analysisIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to save analysis",
			}, http.StatusInternalServerError
		}
		return laboratory, laboratory.LabAnalyses[analysisIndx], http.StatusOK
	})
}

// DeleteLabAnalysis - Delete laboratory analysis by ID
func (t *implLabAnalysesAPI) DeleteLabAnalysis(ctx *gin.Context) {
	updateLaboratoryFunc(ctx, func(c *gin.Context, laboratory *Laboratory) (*Laboratory, interface{}, int) {
		analysisId := ctx.Param("analysisId")

		if analysisId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Analysis ID is required",
			}, http.StatusBadRequest
		}

		analysisIndx := slices.IndexFunc(laboratory.LabAnalyses, func(la LaboratoryAnalysis) bool {
			return analysisId == la.Id
		})

		if analysisIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Analysis not found",
			}, http.StatusNotFound
		}

		laboratory.LabAnalyses = append(laboratory.LabAnalyses[:analysisIndx], laboratory.LabAnalyses[analysisIndx+1:]...)

		return laboratory, nil, http.StatusNoContent
	})
}

// GetLabAnalysis - Get all laboratory analyses
func (t *implLabAnalysesAPI) GetLabAnalysis(ctx *gin.Context) {
	updateLaboratoryFunc(ctx, func(c *gin.Context, laboratory *Laboratory) (*Laboratory, interface{}, int) {
		result := laboratory.LabAnalyses
		if result == nil {
			result = []LaboratoryAnalysis{}
		}
		// return nil lab - no need to update it in db
		return nil, result, http.StatusOK
	})
}

// GetLabAnalysisById - Get laboratory analysis by ID
func (t *implLabAnalysesAPI) GetLabAnalysisById(ctx *gin.Context) {
	updateLaboratoryFunc(ctx, func(c *gin.Context, laboratory *Laboratory) (*Laboratory, interface{}, int) {
		analysisId := ctx.Param("analysisId")

		if analysisId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Analysis ID is required",
			}, http.StatusBadRequest
		}

		analysisIndx := slices.IndexFunc(laboratory.LabAnalyses, func(waiting LaboratoryAnalysis) bool {
			return analysisId == waiting.Id
		})

		if analysisIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Analysis not found",
			}, http.StatusNotFound
		}

		// return nil lab - no need to update it in db
		return nil, laboratory.LabAnalyses[analysisIndx], http.StatusOK
	})
}

// UpdateLabAnalysis - Update laboratory analysis by ID
func (t *implLabAnalysesAPI) UpdateLabAnalysis(ctx *gin.Context) {
	updateLaboratoryFunc(ctx, func(c *gin.Context, laboratory *Laboratory) (*Laboratory, interface{}, int) {
		var analysis LaboratoryAnalysis

		if err := c.ShouldBindJSON(&analysis); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		analysisId := ctx.Param("analysisId")

		if analysisId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Analysis ID is required",
			}, http.StatusBadRequest
		}

		analysisIndx := slices.IndexFunc(laboratory.LabAnalyses, func(waiting LaboratoryAnalysis) bool {
			return analysisId == waiting.Id
		})

		if analysisIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Analysis not found",
			}, http.StatusNotFound
		}

		if analysis.Id != "" {
			laboratory.LabAnalyses[analysisIndx].Id = analysis.Id
		}

		if analysis.Name != "" {
			laboratory.LabAnalyses[analysisIndx].Name = analysis.Name
		}

		if analysis.Analyses != nil {
			laboratory.LabAnalyses[analysisIndx].Analyses = analysis.Analyses
		}

		if analysis.Results != "" {
			laboratory.LabAnalyses[analysisIndx].Results = analysis.Results
		}

		if analysis.Status != "" {
			laboratory.LabAnalyses[analysisIndx].Status = analysis.Status
		}

		return laboratory, laboratory.LabAnalyses[analysisIndx], http.StatusOK
	})
}
