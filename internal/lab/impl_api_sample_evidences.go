package lab

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Copy following section to separate file, uncomment, and implement accordingly
// CreateSampleEvidence - Create a new sample evidence
func (t *implSampleEvidencesAPI) CreateSampleEvidence(ctx *gin.Context) {
	updateLaboratoryFunc(ctx, func(c *gin.Context, laboratory *Laboratory) (*Laboratory, interface{}, int) {
		var sampleEvidence SampleEvidence

		if err := c.ShouldBindJSON(&sampleEvidence); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		if sampleEvidence.Id == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Sample ID is required",
			}, http.StatusBadRequest
		}

		if sampleEvidence.Id == "" || sampleEvidence.Id == "@new" {
			sampleEvidence.Id = uuid.NewString()
		}

		conflictIndx := slices.IndexFunc(laboratory.SampleEvidences, func(se SampleEvidence) bool {
			return sampleEvidence.Id == se.Id
		})

		if conflictIndx >= 0 {
			return nil, gin.H{
				"status":  http.StatusConflict,
				"message": "Sample already exists",
			}, http.StatusConflict
		}

		laboratory.SampleEvidences = append(laboratory.SampleEvidences, sampleEvidence)
		sampleIndx := slices.IndexFunc(laboratory.SampleEvidences, func(se SampleEvidence) bool {
			return sampleEvidence.Id == se.Id
		})
		if sampleIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to save sample",
			}, http.StatusInternalServerError
		}
		return laboratory, laboratory.SampleEvidences[sampleIndx], http.StatusOK
	})
}

// DeleteSampleEvidence - Delete sample evidence by ID
func (t *implSampleEvidencesAPI) DeleteSampleEvidence(ctx *gin.Context) {
	updateLaboratoryFunc(ctx, func(c *gin.Context, laboratory *Laboratory) (*Laboratory, interface{}, int) {
		evidenceId := ctx.Param("evidenceId")

		if evidenceId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Sample ID is required",
			}, http.StatusBadRequest
		}

		sampleIndex := slices.IndexFunc(laboratory.SampleEvidences, func(se SampleEvidence) bool {
			return evidenceId == se.Id
		})

		if sampleIndex < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Sample not found",
			}, http.StatusNotFound
		}

		laboratory.SampleEvidences = append(laboratory.SampleEvidences[:sampleIndex], laboratory.SampleEvidences[sampleIndex+1:]...)

		return laboratory, nil, http.StatusNoContent
	})
}

// GetSampleEvidence - Get all sample evidences
func (t *implSampleEvidencesAPI) GetSampleEvidence(ctx *gin.Context) {
	updateLaboratoryFunc(ctx, func(c *gin.Context, laboratory *Laboratory) (*Laboratory, interface{}, int) {
		result := laboratory.SampleEvidences
		if result == nil {
			result = []SampleEvidence{}
		}
		// return nil ambulance - no need to update it in db
		return nil, result, http.StatusOK
	})
}

// GetSampleEvidenceById - Get sample evidence by ID
func (t *implSampleEvidencesAPI) GetSampleEvidenceById(ctx *gin.Context) {
	updateLaboratoryFunc(ctx, func(c *gin.Context, laboratory *Laboratory) (*Laboratory, interface{}, int) {
		evidenceId := ctx.Param("evidenceId")

		if evidenceId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Sample ID is required",
			}, http.StatusBadRequest
		}

		sampleIndex := slices.IndexFunc(laboratory.SampleEvidences, func(se SampleEvidence) bool {
			return evidenceId == se.Id
		})

		if sampleIndex < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Sample not found",
			}, http.StatusNotFound
		}

		// return nil ambulance - no need to update it in db
		return nil, laboratory.SampleEvidences[sampleIndex], http.StatusOK
	})
}

// UpdateSampleEvidence - Update sample evidence by ID
func (t *implSampleEvidencesAPI) UpdateSampleEvidence(ctx *gin.Context) {
	updateLaboratoryFunc(ctx, func(c *gin.Context, laboratory *Laboratory) (*Laboratory, interface{}, int) {
		var sample SampleEvidence

		if err := c.ShouldBindJSON(&sample); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		sampleId := ctx.Param("evidenceId")

		if sampleId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Sample ID is required",
			}, http.StatusBadRequest
		}

		sampleIndex := slices.IndexFunc(laboratory.SampleEvidences, func(se SampleEvidence) bool {
			return sampleId == se.Id
		})

		if sampleIndex < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Sample not found",
			}, http.StatusNotFound
		}

		if sample.Id != "" {
			laboratory.SampleEvidences[sampleIndex].Id = sample.Id
		}

		return laboratory, laboratory.SampleEvidences[sampleIndex], http.StatusOK
	})
}
