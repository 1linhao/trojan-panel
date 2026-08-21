package dto

type TrafficRankDto struct {
	Period string `json:"period" form:"period" validate:"omitempty,oneof=total month day"`
}

type ServerTrafficUsageDto struct {
	BaseDto
	Period       string `json:"period" form:"period" validate:"required,oneof=total year month day"`
	NodeServerId *uint  `json:"nodeServerId" form:"nodeServerId" validate:"omitempty,gte=0"`
}
