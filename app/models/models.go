package models

import (
	"lpms/app/models/internal/common"
	"lpms/app/models/internal/reserve"
	"lpms/app/models/internal/user"
)

type (
	Base         = common.Base
	User         = user.User
	ReservePro   = reserve.ReservePro
	InvestDetail = reserve.InvestDetail
)
