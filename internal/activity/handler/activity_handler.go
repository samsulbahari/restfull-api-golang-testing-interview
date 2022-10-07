package handler

import (
	"golangtesting/internal/domain"

	"github.com/gin-gonic/gin"
)

type ActivityService interface {
	GetdataService() ([]domain.Activity, int, error)
	GetdataByidService(id string) (domain.Activity, int, error)
	CreateDataService(domain.Activity) (domain.Activity, int, error)
	DeleteDataService(id string) (domain.Activity, int, error)
	UpdateService(activities domain.Activity, id string) (domain.Activity, int, error)
}

type ActivityHandler struct {
	ActivityHan ActivityService
}

func NewActivityHandler(as ActivityService) *ActivityHandler {
	return &ActivityHandler{as}
}

func (ah ActivityHandler) Getdata(ctx *gin.Context) {
	Activity, code, err := ah.ActivityHan.GetdataService()
	if err != nil {
		ctx.JSON(code, gin.H{
			"status":  "Failed",
			"message": err,
		})
		return
	}

	ctx.JSON(code, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Activity,
	})

}

func (ah ActivityHandler) GetdataById(ctx *gin.Context) {
	id := ctx.Param("id")
	Activity, code, err := ah.ActivityHan.GetdataByidService(id)
	if err != nil {
		ctx.JSON(code, gin.H{
			"status":  "Not Found",
			"message": err.Error(),
			"data":    err,
		})
		return
	}

	ctx.JSON(code, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Activity,
	})
}

func (ah ActivityHandler) Createdata(ctx *gin.Context) {
	var activites domain.Activity
	err := ctx.ShouldBindJSON(&activites)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": "title cannot be null",
			"data":    nil,
		})
		return
	}
	Activity, code, err := ah.ActivityHan.CreateDataService(activites)
	if err != nil {
		ctx.JSON(code, gin.H{
			"status":  "Failed",
			"message": err.Error(),
			"data":    err,
		})
		return
	}
	ctx.JSON(code, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Activity,
	})

}

func (ah ActivityHandler) DeleteData(ctx *gin.Context) {
	id := ctx.Param("id")
	_, code, err := ah.ActivityHan.DeleteDataService(id)
	if err != nil {
		if err.Error() == "Success" {
			ctx.JSON(code, gin.H{
				"status":  "Success",
				"message": err.Error(),
				"data":    err,
			})
			return
		} else {
			ctx.JSON(code, gin.H{
				"status":  "Not Found",
				"message": err.Error(),
				"data":    err,
			})
		}

	}

}

func (ah ActivityHandler) UpdateData(ctx *gin.Context) {
	var activites domain.Activity
	err := ctx.ShouldBindJSON(&activites)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": "title cannot be null",
			"data":    nil,
		})
		return
	}
	id := ctx.Param("id")
	Activity, code, err := ah.ActivityHan.UpdateService(activites, id)
	if err != nil {
		ctx.JSON(code, gin.H{
			"status":  "Not Found",
			"message": err.Error(),
			"data":    err,
		})
		return
	}
	ctx.JSON(code, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Activity,
	})

}
