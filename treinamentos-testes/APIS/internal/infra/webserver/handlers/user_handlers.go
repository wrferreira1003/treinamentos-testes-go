package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/wrferreira1003/api-server-go/internal/dto"
	"github.com/wrferreira1003/api-server-go/internal/entity"
	"github.com/wrferreira1003/api-server-go/internal/infra/database"

	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB       database.UserInterface
	jwt          *jwtauth.JWTAuth // Usando o pacote jwtauth do go-chi para autenticação
	jwtExpiresIn int              // Tempo de expiração do token
}

// NewUserHandler cria um novo handler para usuários, serve para criar, atualizar, deletar e buscar usuários
func NewUserHandler(userDB database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:       userDB,
		jwt:          jwt,
		jwtExpiresIn: jwtExpiresIn,
	}
}

// GetJWT gera um token JWT para o usuário
// @Summary Get a user JWT
// @Description Get a user JWT
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.GetJWTInput true "use request"
// @Success 200 {object} dto.GetJWTOutput
// @Failure 404
// @Failure 500
// @Router /users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput
	// Serializando o corpo da requisição para o tipo GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// Se der erro, retorna um status 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verificando se o usuário existe
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized) // status de erro de autenticação
		return
	}

	// Verificando se a senha está correta
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Criando o token
	_, tokenString, _ := h.jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(), // ID do usuário
		"exp": time.Now().Add(time.Second * time.Duration(h.jwtExpiresIn)).Unix(),
	})

	// Criando a serialização do token, para enviar para o cliente
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json") // Setando o header para o cliente entender que é um json
	w.WriteHeader(http.StatusOK)                       // Setando o status para 200
	json.NewEncoder(w).Encode(accessToken)             // Enviando o token para o cliente

}

// Create user godoc
// @Summary Create User
// @Description Create User
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserInput true "use request"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput // dto é um pacote que contém as informações que o usuário vai enviar
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// criando um novo usuário, passando pelo coração do sistema
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// criando o usuário no banco de dados
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
