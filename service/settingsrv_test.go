// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package service

import (
	"testing"

	"github.com/b3log/pipe/model"
)

func TestGetSetting(t *testing.T) {
	setting := Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, 1)
	if nil == setting {
		t.Errorf("setting is nil")

		return
	}

	if "Pipe 示例" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "Pipe 示例", setting.Value)
	}
}

func TestGetAllSettings(t *testing.T) {
	settings := Setting.GetAllSettings(1)
	if 27 != len(settings) {
		t.Errorf("expected is [%d], actual is [%d]", 27, len(settings))
	}
}

func TestGetCategorySettings(t *testing.T) {
	basicSettings := Setting.GetCategorySettings(1, model.SettingCategoryBasic)
	if 10 != len(basicSettings) {
		t.Errorf("expected is [%d], actual is [%d]", 10, len(basicSettings))
	}
}

func TestGetSettings(t *testing.T) {
	settings := Setting.GetSettings(1, model.SettingCategoryBasic,
		[]string{model.SettingNameBasicBlogTitle, model.SettingNameBasicBlogSubtitle})
	if nil == settings {
		t.Errorf("settings is nil")

		return
	}
	if 1 > len(settings) {
		t.Errorf("settings is empty")

		return
	}

	if "Pipe 示例" != settings[model.SettingNameBasicBlogTitle].Value {
		t.Errorf("expected is [%s], actual is [%s]", "Pipe 示例", settings[model.SettingNameBasicBlogTitle].Value)
	}
}

func TestUpdateSettings(t *testing.T) {
	settings := Setting.GetSettings(1, model.SettingCategoryBasic,
		[]string{model.SettingNameBasicBlogTitle, model.SettingNameBasicBlogSubtitle})
	settings[model.SettingNameBasicBlogTitle].Value = "更新后的标题"
	basics := []*model.Setting{}
	for _, setting := range settings {
		basics = append(basics, setting)
	}
	if err := Setting.UpdateSettings(model.SettingCategoryBasic, basics); nil != err {
		t.Errorf("updates settings failed: " + err.Error())

		return
	}

	settings = Setting.GetSettings(1, model.SettingCategoryBasic,
		[]string{model.SettingNameBasicBlogTitle, model.SettingNameBasicBlogSubtitle})
	if "更新后的标题" != settings[model.SettingNameBasicBlogTitle].Value {
		t.Errorf("expected is [%s], actual is [%s]", "更新后的标题", settings[model.SettingNameBasicBlogTitle].Value)
	}
}
