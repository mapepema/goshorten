package secure_test

import (
	"goshorten/pkg/utl/middleware/secure"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSecureCors(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
		headers      map[string]string
	}{
		{
			description:  "get HTTP status 200",
			route:        "/hello",
			expectedCode: 200,
			headers: map[string]string{
				// "Access-Control-Max-Age": "86400", // default not passed
				//"Access-Control-Allow-Methods":     "POST,GET,PUT,DELETE,PATCH,HEAD", // default not passed
				"Access-Control-Expose-Headers": "Content-Length",
				//"Access-Control-Allow-Credentials": "true", // default not passed
			},
		},
	}

	app := fiber.New()
	app.Use(secure.CORS())

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)

		resp, _ := app.Test(req, 1)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

		for header, attended := range test.headers {
			t.Log(header)
			t.Log(resp.Header)
			t.Log(resp.Header.Get("Access-Control-Allow-Origin"))
			assert.Equal(t, resp.Header.Get(header), attended)
		}
	}
}
