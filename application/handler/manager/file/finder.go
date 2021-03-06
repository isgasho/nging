package file

import "github.com/webx-top/echo"

func Finder(ctx echo.Context) error {
	err := List(ctx, ``, 0)
	multiple := ctx.Formx(`multiple`).Bool()
	ctx.Set(`dialog`, true)
	ctx.Set(`multiple`, multiple)
	partial := ctx.Formx(`partial`).Bool()
	if partial {
		return ctx.Render(`manager/file/list.main.content`, err)
	}
	return ctx.Render(`manager/file/finder`, err)
}
