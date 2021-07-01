package dto

import (
	"github.com/gin-gonic/gin"
	"go_gateway_project/public"
)


// 输入结构
type ServiceListInput struct {
	Info string `json:"info" form:"info" comment:"关键字" example:"" validate:""`
	PageNo int `json:"page_no" form:"page_no" comment:"页数" example:"1" validate:"required"`
	PageSize int `json:"page_size" form:"page_size" comment:"每页条数" example:"20" validate:"required"`
}

func (param *ServiceListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

// 输出结构
type ServiceListItemOutput struct {
	Id int64 `json:"id" form:"id" comment:"id" example:"" validate:""`
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名称" example:"" validate:""`
	ServiceDesc string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"" validate:""`
	LoadType string `json:"load_type" form:"load_type" comment:"加载类型" example:"" validate:""`
	ServiceAddr string `json:"service_addr" form:"service_addr" comment:"加载类型" example:"" validate:""`
	Qps string `json:"qps" form:"qps" comment:"每秒请求数" example:"" validate:""`
	Qpd string `json:"qpd" form:"qpd" comment:"每秒返回数" example:"" validate:""`
	TotalNode string `json:"total_node" form:"total_node" comment:"总节点数" example:"" validate:""`
}

type ServiceListOutput struct {
	Total int64 `json:"total" form:"total" comment:"总数" example:"" validate:""`
	List []ServiceListItemOutput `json:"list" form:"list" comment:"列表" example:"" validate:""`
}

