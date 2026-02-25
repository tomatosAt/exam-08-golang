package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomatosAt/exam-08-golang/util"
)

func (h *Handler) GetAllExamHandler(ctx *fiber.Ctx) error {
	res, err := h.svc.GetAllExamSVC(ctx.Context())
	if err != nil {
		return util.HttpError(ctx, 400, err.Error())
	}
	return util.HttpSuccess(ctx, 200, res)
}
