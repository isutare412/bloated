package transaction

import "context"

type Context struct {
	Ctx      context.Context
	Commit   func() error
	Rollback func() error
}

func (c *Context) RollbackIfError(err error) error {
	if err == nil {
		return nil
	}
	return c.Rollback()
}
