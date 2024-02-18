package api

import (
	"net/http"

	db "github.com/Tomlord1122/todo-in-go/db/sqlc"
	"github.com/gin-gonic/gin"
)

// createTodo, listTodos, getTodo, updateTodo, deleteTodo
type CreateTodoParams struct {
	Owner       string `json:"owner"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// form and uri 是gin框架的参数绑定
// form: 通常

type CreateTodoRequest struct {
	Owner       string `json:"owner" required:"true"`
	Title       string `json:"title" required:"true"`
	Category    string `json:"category" required:"true"`
	Description string `json:"description" required:"true"`
	Completed   bool   `json:"completed" required:"true"`
}

type listTodosRequest struct {
	PageID   int32 `form:"page_id" binding:"required"`
	PageSize int32 `form:"page_size" binding:"required"`
}

type getTodoRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateTodoRequest struct {
	ID          int64  `uri:"id" binding:"required,min=1"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type deleteTodoRequest struct {
	id int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) createTodo(ctx *gin.Context) {
	var req CreateTodoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTodoParams{
		Owner:       req.Owner,
		Title:       req.Title,
		Category:    req.Category,
		Description: req.Description,
		Completed:   req.Completed,
	}

	todo, err := server.query.CreateTodo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, todo) // 200
}

func (server *Server) listTodos(ctx *gin.Context) {
	var req listTodosRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTodosParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	todos, err := server.query.ListTodos(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (server *Server) getTodo(ctx *gin.Context) {
	var req getTodoRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	todo, err := server.query.GetTodo(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func (server *Server) updateTodo(ctx *gin.Context) {
	var req updateTodoRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateTodoParams{
		ID:          req.ID,
		Title:       req.Title,
		Category:    req.Category,
		Description: req.Description,
		Completed:   req.Completed,
	}

	todo, err := server.query.UpdateTodo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func (server *Server) deleteTodo(ctx *gin.Context) {
	var req deleteTodoRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.query.DeleteTodo(ctx, req.id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
