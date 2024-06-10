package metadata

import (
	"sync"
	"time"
)

type IOperator interface {
	GetUserID() uint64                      // 获取用户ID
	GetEmpID() uint64                       // 获取用户员工id
	GetUserName() string                    // 获取用户名称
	GetTenantManagementID() uint64          // 获取用户名称
	GetBizUnitID() uint64                   // 获取用户往来单位ID
	GetDefaultSaleSystemID() uint64         // 获取默认营销体系ID
	GetDepartmentID() uint64                // 获取用户所属部门ID
	IsSaasCompany() bool                    // 所属公司是否为saas组织
	IsAvailable() bool                      // 用户是否可用
	GetToken() string                       // 获取用户登录token
	GetSubDepartmentID() []uint64           // 获取用户所在部门的子部门
	GetRoleAccessIDs() []uint64             // 获取用户所属角色权限列表
	GetLoginTime() time.Time                // 获取登录时间
	GetMenuIDs() (r []uint64)               // 获取菜单ID
	GetResourceRouterNames() (r []string)   // 获取前端路由名称
	GetButtonCodes() (r []string)           // 获取前端按钮编号
	GetMPResourceRouterNames() (r []string) // 获取前端内部商城路由名称
	GetMPButtonCodes() (r []string)         // 获取前端内部商城按钮编号
	GetWarehouseIDs() (r []uint64)          // 获取仓库权限
	GetSaleSystemIDs() (r []uint64)         // 获取营销体系权限
	GetBizUnitIDs() (r []uint64)            // 获取往来单位权限
	IsAllowUpdateOrder() (r bool)           // 是否允许更新别人的单据
	IsAllowCancelOther() (r bool)           // 是否允许作废别人的单据
	IsAllowAuditSelf() (r bool)             // 是否允许审核自己的单据
}

type IMallOperator interface {
	IOperator
	GetWeChatOpenUserID() uint64 // 获取用户openid
	GetAppid() string            // 获取用户所属小程序appid
}

var UserInfo sync.Map
