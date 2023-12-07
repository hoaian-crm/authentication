package role_dto

type UpdateRoleUri struct {
	RoleId int64 `uri:"roleId" binding:"must_found=roles"`
}
