package vo

import (
	"lpms/app/models"

	"github.com/goccy/go-json"
)

type WindowsReq struct {
	// 储备库填报
	ReserveSetting string `json:"reserve_setting"`
	// 项目进度填报
	ProgressSetting string `json:"progress_setting"`
	// 项目计划填报
	ProPlanSetting string `json:"pro_plan_setting"`
}

func (w *WindowsReq) ToModel(openID string) *models.WindowSetting {
	return &models.WindowSetting{
		ReserveSetting:  json.RawMessage([]byte(w.ReserveSetting)),
		ProgressSetting: json.RawMessage([]byte(w.ProgressSetting)),
		ProPlanSetting:  json.RawMessage([]byte(w.ProPlanSetting)),
		Base: models.Base{
			UpdateBy: openID,
			CreateBy: openID,
		},
	}
}

type WindowsResponse struct {
	ReserveSetting  string `json:"reserve_setting"`
	ProgressSetting string `json:"progress_setting"`
	ProPlanSetting  string `json:"pro_plan_setting"`
}

func NewWindowsResponse(w *models.WindowSetting) *WindowsResponse {
	return &WindowsResponse{
		ReserveSetting:  string(w.ReserveSetting),
		ProgressSetting: string(w.ProgressSetting),
		ProPlanSetting:  string(w.ProPlanSetting),
	}
}

type WindowsUpdateReq struct {
	// 储备库填报
	ReserveSetting string `json:"reserve_setting"`
	// 项目进度填报
	ProgressSetting string `json:"progress_setting"`
	// 项目计划填报
	ProPlanSetting string `json:"pro_plan_setting"`
}

func (w *WindowsUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"reserve_setting":  w.ReserveSetting,
		"progress_setting": w.ProgressSetting,
		"pro_plan_setting": w.ProPlanSetting,
		"update_by":        openID,
	}
}
