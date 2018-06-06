package model

import "HearthStoneDatabase/entity"

type Hero struct {
	entity.Hero
	SkillName string
	SkillImageSrc string
	ProfessionName string
}