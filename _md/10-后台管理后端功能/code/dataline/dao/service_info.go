package dao

import (
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"go_gateway_project/dto"
	"go_gateway_project/public"
	"time"
)

type ServiceInfo struct {
	Id        int64     `json:"id" gorm:"primary_key" description:"自增主键"`
	LoadType  int    	`json:"load_type" gorm:"column:load_type" description:"负载类型 0=http 1=tcp 2=grpc"`
	ServiceName string  `json:"service_name" gorm:"column:service_name" description:"服务名称"`
	ServiceDesc string  `json:"service_desc" gorm:"column:service_desc" description:"服务描述"`
	IsDelete  int64     `json:"is_delete" gorm:"column:is_delete" description:"是否删除 0=否 1=是"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
}

func (t *ServiceInfo) TableName() string {
	return "gateway_service_info"
}

// 分页查询
func (t *ServiceInfo) PageList(ctx *gin.Context, tx *gorm.DB, param *dto.ServiceListInput) ([]ServiceInfo, int64, error) {
	total := int64(0)
	offset := param.PageSize*(param.PageNo-1)
	list:=[]ServiceInfo{}

	// 查询数据
	query:=tx.SetCtx(public.GetTraceContext(ctx))
	query = query.Table(t.TableName()).Where("is_delete=0")

	if param.Info!="" {
		query = query.Where("service_name like %%?% or service_desc like %%?%", param.Info)
	}

	if err:=query.Limit(param.PageSize).Offset(offset).Find(&list).Error; err==gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	// 计算总数
	query.Count(&total)
	return list, total, nil
}

// params.Username 取得管理员信息 admininfo
func (t *ServiceInfo) Find(c *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceInfo, error) {
	out := &ServiceInfo{}
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	// .Where("id = ?", id)

	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *ServiceInfo) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error
}
