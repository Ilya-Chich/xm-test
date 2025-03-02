package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"xm-test-ilya-chicherin/internal/controllers/v1/view"

	"xm-test-ilya-chicherin/entity/company"
)

func (c *CompanyController) GetCompany(ctx *gin.Context) {
	companyID := ctx.Param("id")
	cmp, err := c.CompanyUC.GetCompany(ctx, companyID)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, cmp)
}

func (c *CompanyController) DeleteCompany(ctx *gin.Context) {
	companyID := ctx.Param("id")

	err := c.CompanyUC.DeleteCompany(ctx, companyID)
	if err != nil {
		log.Printf("failed to delete company: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *CompanyController) CreateCompany(ctx *gin.Context) {
	var req company.Company
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err := c.CompanyUC.CreateCompany(ctx, req)
	if err != nil {
		log.Printf("failed to create company: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *CompanyController) PatchCompany(ctx *gin.Context) {
	companyID := ctx.Param("id")

	var req view.PatchCompanyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	updates := make(map[string]interface{})
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.AmountOfEmployees != nil {
		updates["amount_of_employees"] = *req.AmountOfEmployees
	}
	if req.Registered != nil {
		updates["registered"] = *req.Registered
	}
	if req.Type != nil {
		updates["type"] = *req.Type
	}

	if len(updates) == 0 {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err := c.CompanyUC.UpdateCompany(ctx, companyID, updates)
	if err != nil {
		log.Printf("failed to update company: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}
