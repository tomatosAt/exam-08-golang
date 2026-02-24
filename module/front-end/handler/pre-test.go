package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomatosAt/exam-08-golang/util"
)

func (h *Handler) CheckHandler(ctx *fiber.Ctx) error {
	return util.HttpSuccess(ctx, 200, "Check Handler")
}
