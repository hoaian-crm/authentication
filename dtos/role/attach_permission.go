package role_dto

type AttachRolePermission struct {
	PermissionId int64 `uri:"permissionId" binding:"must_found=permissions"`
	RoleId       int64 `uri:"roleId" binding:"must_found=roles"`
}

type AttachRolePatchPermisison struct {
	RoleId        int64   `uri:"roleId"`
	PermissionIds []int64 `json:"permissions"`
}
