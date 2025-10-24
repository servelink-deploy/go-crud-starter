package handlers

import (
	"go-crud-starter/models"
	"go-crud-starter/repository"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		repo: repository.NewUserRepository(),
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Erreur de validation: " + err.Error(),
		})
		return
	}

	exists, err := h.repo.EmailExists(req.Email, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erreur lors de la vérification de l'email",
		})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, models.ErrorResponse{
			Error: "Un utilisateur avec cet email existe déjà",
		})
		return
	}

	user, err := h.repo.Create(&req)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, models.ErrorResponse{
				Error: "Un utilisateur avec cet email existe déjà",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erreur lors de la création de l'utilisateur",
		})
		return
	}

	c.JSON(http.StatusCreated, models.SuccessResponse{
		Message: "Utilisateur créé avec succès",
		Data:    user,
	})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	result, err := h.repo.FindAll(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erreur lors de la récupération des utilisateurs",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "ID invalide",
		})
		return
	}

	user, err := h.repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erreur lors de la récupération de l'utilisateur",
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Utilisateur non trouvé",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) SearchUsers(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Paramètre de recherche 'q' requis",
		})
		return
	}

	users, err := h.repo.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erreur lors de la recherche",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": users,
		"count":   len(users),
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "ID invalide",
		})
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Erreur de validation: " + err.Error(),
		})
		return
	}

	existingUser, err := h.repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erreur lors de la vérification de l'utilisateur",
		})
		return
	}

	if existingUser == nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Utilisateur non trouvé",
		})
		return
	}

	if req.Email != nil && *req.Email != existingUser.Email {
		exists, err := h.repo.EmailExists(*req.Email, &id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: "Erreur lors de la vérification de l'email",
			})
			return
		}

		if exists {
			c.JSON(http.StatusConflict, models.ErrorResponse{
				Error: "Un utilisateur avec cet email existe déjà",
			})
			return
		}
	}

	user, err := h.repo.Update(id, &req)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, models.ErrorResponse{
				Error: "Un utilisateur avec cet email existe déjà",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erreur lors de la mise à jour de l'utilisateur",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Utilisateur mis à jour avec succès",
		Data:    user,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "ID invalide",
		})
		return
	}

	deleted, err := h.repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erreur lors de la suppression de l'utilisateur",
		})
		return
	}

	if !deleted {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Utilisateur non trouvé",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Utilisateur supprimé avec succès",
	})
}
