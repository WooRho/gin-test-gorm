package metadata

import (
	"context"
	"strconv"
)

const (
	UserName           = "user_name"
	UserID             = "user_id"
	CompanyID          = "company_id"
	DepartmentID       = "department_id"
	BizUnitID          = "biz_unit_id"
	SaleSystemID       = "sale_system_id"
	TenantManagementID = "tenant_management_id"

	LoginInfo          = "login_info"
	DataSeparate       = "data_separate"
	IsAllowUpdateOrder = "is_allow_update_order"
	IsAllowCancelOther = "is_allow_cancel_other"
	IsAllowAuditSelf   = "is_allow_audit_self"
)

func SetUserNameToIncoming(ctx context.Context, name string) context.Context {
	return SetMDToIncoming(ctx, UserName, name)
}

func SetUserIDToIncoming(ctx context.Context, uid uint64) context.Context {
	return SetMDToIncoming(ctx, UserID, strconv.FormatUint(uid, 10))
}

func GetUserID(ctx context.Context) uint64 {
	str := GetMD(ctx, UserID)
	if str == "" {
		return 0
	}
	userID, err := strconv.ParseUint(str, 64, 20)
	if err != nil {
		return 0
	}
	return userID
}
