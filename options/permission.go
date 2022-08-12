package options

import (
	"github.com/heinrian/permify-gorm/utils"
)

// PermissionOption represents options when fetching permissions.
type PermissionOption struct {
	Pagination *utils.Pagination
}
