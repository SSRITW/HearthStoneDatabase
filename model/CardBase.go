package model

import "HearthStoneDatabase/entity"

type CardBase struct {
	entity.CardBase
	CardTypeName string
	CardPackageName string
	CardPackageBackImage string
	ProfessionCardFront string
}
