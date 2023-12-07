package role_dto

type DeleteRoleUri struct {
	RoleId int64 `uri:"roleId" binding:"must_found=roles"`
}
