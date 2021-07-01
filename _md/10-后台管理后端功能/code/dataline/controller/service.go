package controller

import (
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"go_gateway_project/dao"
	"go_gateway_project/dto"
	"go_gateway_project/middleware"
)

type ServiceController struct{}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("/service_list", service.ServiceList)
}

// ServiceList godoc
// @Summary 服务列表
// @Description 服务列表
// @Tags 服务管理
// @ID /service/service_list
// @Accept  json
// @Produce  json
// @Param info query string false "关键字"
// @Param page_no query int true "页数"
// @Param page_size query int true "每页条数"
// @Success 200 {object} middleware.Response{data=dto.ServiceListItemOutput} "success"
// @Router /service/service_list [get]
func (service *ServiceController) ServiceList(ctx *gin.Context) {
	// 校验参数
	params := &dto.ServiceListInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}

	// 获取数据库链接
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}

	// 数据库链接查询数据
	serviceInfo := &dao.ServiceInfo{}
	list, total, err := serviceInfo.PageList(ctx, tx, params)
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}

	// 格式化查询数据输出
	outList := []dto.ServiceListItemOutput{}
	for _,listItem:=range  list {
		outItem:=dto.ServiceListItemOutput{
			Id: listItem.Id,
			ServiceName: listItem.ServiceName,
			ServiceDesc: listItem.ServiceDesc,
		}

		outList = append(outList, outItem)
	}

	out := &dto.ServiceListOutput{
		Total:total,
		List:outList,
	}

	fmt.Println(out)
	middleware.ResponseSuccess(ctx, out)
}