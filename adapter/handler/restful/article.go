package restful

import (
	"net/http"
	"realworld-go-fiber/core/domain"
	"realworld-go-fiber/core/port"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type articleHandler struct {
	articleUc port.ArticleUsecase
	validator *validator.Validate
}

func NewArticleHandler(au port.ArticleUsecase, validator *validator.Validate) *articleHandler {
	return &articleHandler{articleUc: au, validator: validator}
}

type CreateArticle struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Body        string   `json:"body" validate:"required"`
	TagList     []string `json:"tagList"`
}

type CreateArticleRequest struct {
	Article CreateArticle `json:"article"`
}

func (h *articleHandler) CreateArticle(c *fiber.Ctx) error {
	arg, err := getAuthArg(c)
	if err != nil {
		return errorHandler(c, err)
	}

	req := CreateArticleRequest{}
	if err := c.BodyParser(&req); err != nil {
		return errorHandler(c, err)
	}
	if err := h.validator.Struct(&req); err != nil {
		return errorHandler(c, err)
	}

	article, err := h.articleUc.CreateArticle(port.CreateArticleParams{
		Article: domain.Article{
			Title:       req.Article.Title,
			Description: req.Article.Description,
			Body:        req.Article.Body,
			TagNames:    req.Article.TagList,
		},
		AuthArg: arg,
	})
	if err != nil {
		return errorHandler(c, err)
	}

	res := ArticleResponse{Article: serializeArticle(article)}
	return c.Status(http.StatusCreated).JSON(res)
}

func (h *articleHandler) GetArticle(c *fiber.Ctx) error {
	arg, err := getAuthArg(c)
	if err != nil {
		return errorHandler(c, err)
	}

	slug := c.Params("slug")

	article, err := h.articleUc.GetArticle(port.GetArticleParams{
		AuthArg: arg,
		Article: domain.Article{Slug: slug},
	})
	if err != nil {
		return errorHandler(c, err)
	}

	res := ArticleResponse{Article: serializeArticle(article)}
	return c.Status(http.StatusCreated).JSON(res)
}
