package role_dto

type DetachPermission struct {
	PermissionId int64 `uri:"permissionId" binding:"must_found=permissions"`
	RoleId       int64 `uri:"roleId" binding:"must_found=roles"`
}
