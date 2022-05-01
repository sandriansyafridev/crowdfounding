package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/crowdfounding/model/format"
	"github.com/sandriansyafridev/crowdfounding/model/response"
	"github.com/sandriansyafridev/crowdfounding/service"
)

type CampaignController interface {
	GetCampaigns(c *gin.Context)
	GetCampaign(c *gin.Context)
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
