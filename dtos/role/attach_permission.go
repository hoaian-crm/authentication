package role_dto

type AttachRolePermission struct {
	PermissionId uint `uri:"permissionId" binding:"must_found=permissions"`
	RoleId       uint `uri:"roleId" binding:"must_found=roles"`
}
