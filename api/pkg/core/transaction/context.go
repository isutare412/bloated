package transaction

import "context"

type Context struct {
	Ctx      context.Context
	Commit   func() error
	Rollback func() error
}
