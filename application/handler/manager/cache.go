/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package manager

import (
	"github.com/coscms/webcore/library/config"
	"github.com/coscms/webcore/library/modal"
	"github.com/coscms/webcore/library/notice"
	"github.com/webx-top/echo"
)

func ClearCache(ctx echo.Context) error {
	if err := modal.Clear(); err != nil {
		return err
	}
	notice.Clear()
	ctx.FireByNameWithMap(`nging.renderer.cache.clear`, echo.H{`ctx`: ctx})
	return ctx.String(ctx.T(`已经清理完毕`))
}

func ReloadEnv(ctx echo.Context) error {
	if err := config.FromCLI().InitEnviron(true); err != nil {
		return err
	}
	ctx.FireByNameWithMap(`nging.env.reload`, echo.H{`ctx`: ctx})
	return ctx.String(ctx.T(`重载完毕`))
}
