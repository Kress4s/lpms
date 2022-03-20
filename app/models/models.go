package models

import (
	"lpms/app/models/internal/common"
	"lpms/app/models/internal/implement"
	"lpms/app/models/internal/reserve"
	"lpms/app/models/internal/user"
)

type (
	Base           = common.Base
	User           = user.User
	ReservePro     = reserve.ReservePro
	InvestDetail   = reserve.InvestDetail
	ListReservePro = reserve.ListReservePro

	Object = common.Object

	ImplementGov        = implement.ImplementGov
	ImpleIndustry       = implement.ImpleIndustry
	GovProgress         = implement.GovProgress
	ListGovProgressPlan = implement.ListGovProgressPlan
	GovProgressCompare  = implement.GovProgressCompare
)
