package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestExampleHandler testa o handler ExampleHandler para verificar se a resposta é válida e correta.
func TestExampleHandler(t *testing.T) {
	// Setup Gin para rodar os testes
	r := gin.Default()
	r.GET("/example", ExampleHandler)

	// Cria uma requisição de teste
	req, _ := http.NewRequest("GET", "/example", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verifica se o status da resposta é 200
	assert.Equal(t, 200, w.Code)
	// Verifica se a resposta contém a mensagem que estamos esperando receber
	assert.Contains(t, w.Body.String(), "Hello from the handler!")
}
