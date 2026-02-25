package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomatosAt/exam-08-golang/module/front-end/dto"
	"github.com/tomatosAt/exam-08-golang/util"
)

func (h *Handler) GetAllExamHandler(ctx *fiber.Ctx) error {
	res, err := h.svc.GetAllExamSVC(ctx.Context())
	if err != nil {
		return util.HttpError(ctx, 400, err.Error())
	}
	return util.HttpSuccess(ctx, 200, res)
}

func (h *Handler) CreateExamHandler(ctx *fiber.Ctx) error {
	var req dto.CreateExamRequest
	if err := ctx.BodyParser(&req); err != nil {
		return util.HttpError(ctx, 400, err.Error())
	}
	if err := h.validator.Struct(req); err != nil {
		return util.HttpError(ctx, 400, err.Error())
	}
	res, err := h.svc.CreateExamSVC(ctx.Context(), req)
	if err != nil {
		return util.HttpError(ctx, 400, err.Error())
	}
	return util.HttpSuccess(ctx, 200, res)
}
