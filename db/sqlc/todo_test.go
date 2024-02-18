package db

import (
	"context"
	"log"
	"testing"

	"github.com/Tomlord1122/todo-in-go/util"
	"github.com/stretchr/testify/require"
)

func createRandomTodo(t *testing.T) Todo {
	arg := CreateTodoParams{
		Owner:     util.RandomString(4),
		Title:     util.RandomString(10),
		Category:  util.RandomString(3),
		Completed: util.RandomBool(),
	}

	todo, err := testQueries.CreateTodo(context.Background(), arg)
	if err != nil {
		log.Fatal("cannot create random todo:", err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.Equal(t, arg.Owner, todo.Owner)
	require.Equal(t, arg.Title, todo.Title)
	require.Equal(t, arg.Category, todo.Category)
	require.Equal(t, arg.Completed, todo.Completed)

	require.NotZero(t, todo.ID)
	return todo
}

func TestCreateTodo(t *testing.T) {
	createRandomTodo(t)
}

func TestGetTodo(t *testing.T) {
	todo1 := createRandomTodo(t)
	todo2, err := testQueries.GetTodo(context.Background(), todo1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, todo2)

	require.Equal(t, todo1.Owner, todo2.Owner)
	require.Equal(t, todo1.Title, todo2.Title)
	require.Equal(t, todo1.Category, todo2.Category)
	require.Equal(t, todo1.Completed, todo2.Completed)
}

func TestDeleteTodo(t *testing.T) {
	todo1 := createRandomTodo(t)
	err := testQueries.DeleteTodo(context.Background(), todo1.ID)
	require.NoError(t, err)

	todo2, err := testQueries.GetTodo(context.Background(), todo1.ID)
	require.Error(t, err)
	require.Empty(t, todo2)
}

func TestGetsTodo(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTodo(t)
	}

	arg := ListTodosParams{
		Limit:  5,
		Offset: 5,
	}

	todos, err := testQueries.ListTodos(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, todos, 5)

	for _, todo := range todos {
		require.NotEmpty(t, todo)
	}
}

func TestUpdateTodo(t *testing.T) {
	todo1 := createRandomTodo(t)

	arg := UpdateTodoParams{
		ID:          todo1.ID,
		Title:       util.RandomString(10),
		Category:    util.RandomString(3),
		Description: util.RandomString(5),
		Completed:   util.RandomBool(),
	}

	todo2, err := testQueries.UpdateTodo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo2)

	require.Equal(t, todo1.ID, todo2.ID)
	require.Equal(t, arg.Title, todo2.Title)
	require.Equal(t, arg.Category, todo2.Category)
	require.Equal(t, arg.Description, todo2.Description)
	require.Equal(t, arg.Completed, todo2.Completed)
}
