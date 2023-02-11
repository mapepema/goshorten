package secure_test

import (
	"goshorten/pkg/utl/middleware/secure"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)


func TestSecureHaeders(t *testing.T) {
	tests := []struct {
		description string
		route string
		expectedCode int
		headers map[string]string
	} {
		{
			description: "get HTTP status 200",
			route: "/hello",
			expectedCode: 200,
			headers: map[string]string{
				"X-Content-Type-Options": "nosniff",
				"X-DNS-Prefetch-Control": "off",
				"X-Frame-Options": "DENY",
				"Strict-Transport-Security": "max-age=5184000; includeSubDomains",
				"X-Download-Options": "noopen",
				"X-XSS-Protection": "1; mode=block",
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
			assert.Equal(t, req.)
		}
	}


}