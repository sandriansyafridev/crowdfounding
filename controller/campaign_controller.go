package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/crowdfounding/model/dto"
	"github.com/sandriansyafridev/crowdfounding/model/entity"
	"github.com/sandriansyafridev/crowdfounding/model/format"
	"github.com/sandriansyafridev/crowdfounding/model/response"
	"github.com/sandriansyafridev/crowdfounding/service"
)

type CampaignController interface {
	GetCampaigns(c *gin.Context)
	GetCampaign(c *gin.Context)
	CreateCampaign(c *gin.Context)
	UpdateCampaign(c *gin.Context)
}

type CampaignControllerImpl struct {
	service.CampaignService
}

func NewCampaignControllerImpl(campaignService service.CampaignService) *CampaignControllerImpl {
	return &CampaignControllerImpl{
		CampaignService: campaignService,
	}
}

func (campaignController *CampaignControllerImpl) GetCampaigns(c *gin.Context) {

	if queryParamsUserID := c.Query("user_id"); queryParamsUserID == "" {
		if campaigns, err := campaignController.CampaignService.GetCampaigns(); err != nil {
			responseFail := response.ResponseFail("fail fetch campaigns", err)
			c.JSON(http.StatusNotFound, responseFail)
		} else {
			responseSuccess := response.ResponseSuccess("success fetch campaigns", format.ToCampaignsResponse(campaigns))
			c.JSON(http.StatusOK, responseSuccess)
		}
	} else {

		UserID, _ := strconv.Atoi(queryParamsUserID)
		if campaigns, err := campaignController.CampaignService.GetCampaignsByUserID(uint64(UserID)); err != nil {
			responseFail := response.ResponseFail("fail fetch campaigns by user id", err)
			c.JSON(http.StatusNotFound, responseFail)
		} else {
			responseSuccess := response.ResponseSuccess("success fetch campaigns by user id", format.ToCampaignsResponse(campaigns))
			c.JSON(http.StatusOK, responseSuccess)
		}

	}

}

func (campaignController *CampaignControllerImpl) GetCampaign(c *gin.Context) {

	queryCampaignID := c.Param("id")
	if queryCampaignID == "" {
		responseFail := response.ResponseFail("campaign id no params", errors.New("no params"))
		c.JSON(http.StatusNotFound, responseFail)
	} else {
		CampaignID, _ := strconv.Atoi(queryCampaignID)
		if campaign, err := campaignController.CampaignService.GetCampaign(uint64(CampaignID)); err != nil {
			responseFail := response.ResponseFail("fail to fetch campaign", err)
			c.JSON(http.StatusNotFound, responseFail)
		} else {
			responseSuccess := response.ResponseSuccess("success to fetch campaign", format.ToCampaignDetailResponse(campaign))
			c.JSON(http.StatusOK, responseSuccess)
		}
	}

}

func (campaignController *CampaignControllerImpl) CreateCampaign(c *gin.Context) {

	request := dto.CampaignCreateDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		responseFail := response.ResponseFail("bad request", err)
		c.JSON(http.StatusBadRequest, responseFail)
	} else {

		currentUser := c.MustGet("currentUser").(entity.User)
		request.UserID = currentUser.ID
		if campaign, err := campaignController.CampaignService.CreateCampaign(request); err != nil {
			responseFail := response.ResponseFail("fail to create campaign", err)
			c.JSON(http.StatusUnprocessableEntity, responseFail)
		} else {
			responseSuccess := response.ResponseSuccess("campaign created", format.ToCampaignResponse(campaign))
			c.JSON(http.StatusOK, responseSuccess)
		}

	}

}

func (campaignController *CampaignControllerImpl) UpdateCampaign(c *gin.Context) {
	request := dto.CampaignUpdateDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		responseFail := response.ResponseFail("bad request", err)
		c.JSON(http.StatusBadRequest, responseFail)
	} else {

		CampaignID, _ := strconv.Atoi(c.Param("id"))
		currentUser := c.MustGet("currentUser").(entity.User)
		request.ID = uint64(CampaignID)
		request.UserID = currentUser.ID
		if campaign, err := campaignController.CampaignService.UpdateCampaign(request); err != nil {
			responseFail := response.ResponseFail("fail to update campaign", err)
			c.JSON(http.StatusUnprocessableEntity, responseFail)
		} else {
			responseSuccess := response.ResponseSuccess("campaign updated", format.ToCampaignResponse(campaign))
			c.JSON(http.StatusOK, responseSuccess)
		}
	}
}
