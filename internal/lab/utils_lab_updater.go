package lab

import (
	"net/http"

	"github.com/Magum23/wac-lab-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
)

type laboratoryUpdater = func(
	ctx *gin.Context,
	laboratory *Laboratory,
) (updatedLaboratory *Laboratory, responseContent interface{}, status int)

func updateLaboratoryFunc(ctx *gin.Context, updater laboratoryUpdater) {
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

	db, ok := value.(db_service.DbService[Laboratory])
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

	laboratoryId := ctx.Param("laboratoryId")

	laboratory, err := db.FindDocument(ctx, laboratoryId)

	switch err {
	case nil:
		// continue
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Laboratory not found",
				"error":   err.Error(),
			},
		)
		return
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to load laboratory from database",
				"error":   err.Error(),
			})
		return
	}

	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Failed to cast laboratory from database",
				"error":   "Failed to cast laboratory from database",
			})
		return
	}

	updatedLaboratory, responseObject, status := updater(ctx, laboratory)

	if updatedLaboratory != nil {
		err = db.UpdateDocument(ctx, laboratoryId, updatedLaboratory)
	} else {
		err = nil // redundant but for clarity
	}

	switch err {
	case nil:
		if responseObject != nil {
			ctx.JSON(status, responseObject)
		} else {
			ctx.AbortWithStatus(status)
		}
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Laboratory was deleted while processing the request",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to update laboratory in database",
				"error":   err.Error(),
			})
	}

}
