package manager

import (
	"github.com/webx-top/com"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
	"github.com/webx-top/echo/subdomains"

	"github.com/coscms/webcore/dbschema"
	"github.com/coscms/webcore/library/config"
	"github.com/coscms/webcore/model"
	"github.com/coscms/webcore/model/file/storer"
	"github.com/coscms/webcore/registry/settings"
	"github.com/coscms/webcore/registry/upload/driver"
	"github.com/coscms/webcore/registry/upload/driver/local"
	"github.com/coscms/webcore/registry/upload/driver/s3"

	"github.com/webx-top/image"
)

var configDefaults = map[string]map[string]*dbschema.NgingConfig{
	`base`: {
		`storer`: {
			Key:         `storer`,
			Label:       echo.T(`存储引擎`),
			Description: ``,
			Value:       `{"name": "local", "id": ""}`,
			Group:       `base`,
			Type:        `json`,
			Sort:        30,
			Disabled:    `N`,
		},
		`watermark`: {
			Key:         `watermark`,
			Label:       echo.T(`图片水印`),
			Description: ``,
			Value:       `{"watermark": "/public/assets/backend/images/nging-gear.png", "type": "image", "position": 0, "padding": 0, "on": true}`,
			Group:       `base`,
			Type:        `json`,
			Sort:        30,
			Disabled:    `N`,
		},
	},
	`captcha`: {
		`type`: {
			Key:         `type`,
			Label:       echo.T(`验证码类型`),
			Description: ``,
			Value:       ``,
			Group:       `captcha`,
			Type:        `text`,
			Sort:        30,
			Disabled:    `N`,
		},
		`api`: {
			Key:         `api`,
			Label:       echo.T(`第三方验证码接口设置`),
			Description: ``,
			Value:       ``,
			Group:       `captcha`,
			Type:        `json`,
			Sort:        30,
			Disabled:    `N`,
		},
		`go`: {
			Key:         `go`,
			Label:       echo.T(`行为验证码`),
			Description: ``,
			Value:       ``,
			Group:       `captcha`,
			Type:        `json`,
			Sort:        30,
			Disabled:    `N`,
		},
	},
}

var defaultStorer = storer.Info{
	Name: local.Name,
}

func init() {
	// 添加默认配置数据
	for group, configs := range configDefaults {
		settings.AddDefaultConfig(group, configs)
	}
	// 注册配置模板和逻辑
	if index, setting := settings.Get(`base`); index != -1 && setting != nil {
		setting.AddHookGet(func(ctx echo.Context) error {
			ctx.Set(`storerNames`, driver.AllNames())
			m := model.NewCloudStorage(ctx)
			m.ListByOffset(nil, nil, 0, -1)
			ctx.Set(`cloudStorageAccounts`, m.Objects())
			return nil
		})
	}
	settings.Register(&settings.SettingForm{
		Short:    echo.T(`验证码设置`),
		Label:    echo.T(`验证码设置`),
		Group:    `captcha`,
		Tmpl:     []string{`manager/settings/captcha`},
		FootTmpl: []string{`manager/settings/captcha_footer`},
	})
	settings.RegisterDecoder(`base.storer`, func(v *dbschema.NgingConfig, r echo.H) error {
		jsonData := storer.NewInfo()
		if len(v.Value) > 0 {
			com.JSONDecode(com.Str2bytes(v.Value), jsonData)
		}
		r[`ValueObject`] = jsonData
		return nil
	})
	settings.RegisterDecoder(`base.watermark`, func(v *dbschema.NgingConfig, r echo.H) error {
		jsonData := image.NewWatermarkOptions()
		if len(v.Value) > 0 {
			com.JSONDecode(com.Str2bytes(v.Value), jsonData)
		}
		r[`ValueObject`] = jsonData
		return nil
	})
	settings.RegisterEncoder(`base.storer`, func(v *dbschema.NgingConfig, r echo.H) ([]byte, error) {
		cfg := storer.NewInfo().FromStore(r)
		if cfg.Name == local.Name {
			cfg.ID = ``
		} else {
			id := param.AsUint(cfg.ID)
			if id > 0 {
				cfg.Name = s3.Name
			} else {
				cfg.Name = local.Name
			}
		}
		return com.JSONEncode(cfg)
	})
	settings.RegisterEncoder(`base.watermark`, func(v *dbschema.NgingConfig, r echo.H) ([]byte, error) {
		cfg := image.NewWatermarkOptions().FromStore(r)
		return com.JSONEncode(cfg)
	})
	var updateStorer = func(diff config.Diff) error {
		settings, ok := diff.New.(*storer.Info)
		if !ok || settings == nil {
			settings = &defaultStorer
		} else {
			settings.Cloud()
		}
		echo.Set(storer.StorerInfoKey, settings)
		return nil
	}
	config.OnKeySetSettings(`base.storer`, updateStorer)
	var updateBackendURL = func(diff config.Diff) error {
		subdomains.SetBaseURL(`backend`, diff.String())
		return nil
	}
	config.OnKeySetSettings(`base.backendURL`, updateBackendURL)
}
