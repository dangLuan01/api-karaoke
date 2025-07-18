package v1handler

import (
	"net/http"

	v1service "github.com/dangLuan01/karaoke/internal/service/v1"
	"github.com/dangLuan01/karaoke/internal/utils"
	"github.com/gin-gonic/gin"
)

type SongHandler struct {
	service v1service.SongService
}

func NewSongHandler(service v1service.SongService) *SongHandler {
	return &SongHandler{
		service: service,
	}
}

func (sh *SongHandler) GetAll(ctx *gin.Context) {
	data := "U2FsdGVkX1/hStGQ2Nhi5gEu5Xne9/QZrTPoKgWZfuSAP2Gxs+z4RvFnu5djGST8n0++cM/K4l2k94aBMPu5r8NsGHm8uo6uNYBAV5JSnrPQzzh9rimZWRs2vw7qZiTLNh3qo3Ll7Xt6VkXM8MA33VEIIPlMo6vdJD5HFnTzPOfFNN963lEc89Bcg+5D3e66Fb6gs9wf3Sz06BQhp2KDBYGIwsKNV37BSPZCdUG635wYDdpi6buupG7grgIMnlF51Ix+juQNCU9lARVIJy4pAKaZALu4BHzwM+lyfjDZdX5+ShIvFP9o6DqcLdb3HgmPxn+zelkXJ3FzaTp5einU5bhSuuLtJkBJF/N5BQfa3AaAECjoXCRqzNJtDgJLtipvbPt1Trrx59plQZ0NVSsVKHDxsWoERhDejKbrqzMmxoq62wctweWmTy4tMp7Ex5Ua7Jzr32KxIMc98kU+CL0Q/+T9IWwZIDyGFGErwGy5xVG3yMBWay0ZnGY5KPkNQ2sNv7yrHqODPAmKy97u7IO2kbVY326sj2BhvAA4b4K+l8r3T8Irc3NYZS8fhAULQOji1gAGT/wpbhPAWODASgZBzO8u/BgAIxc69TPQYeglsXSmF1c5TV+lp2XmMnnwvRMS4Qv+5Q7tgT+5l0cbK3C8gzSOWhy9pefdrSrmqf9NbeJtf1cj1T4eKkACz52HtwJtX9DikCupcmQdFGhzgyKPdIpQ+cDPmcJtQMeKAQgVTaUGQF2eTlVGeWlESlC3f/uOTqahmQSBUFx1Nw19ZTZJXo6568i8lTuzM9NpERLKmkB/oejKYfREL6bIOSdbRFGAFxmw+8WVUh8B3Rlr9uTiv0gX3nZ53StDV4vKYo6LLN0="
	song, err := sh.service.GetAll(data)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, song)
}