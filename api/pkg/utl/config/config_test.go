package config_test

import (
	"goshorten/pkg/utl/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	cases := []struct {
		name     string
		path     string
		wantData *config.Configuration
		wantErr  bool
	}{
		{
			name:    "Fail on non-existing file",
			path:    "notExists",
			wantErr: true,
		},
		{
			name:    "Fail on wrong file format",
			path:    "testdata/config.invalid.yaml",
			wantErr: true,
		},
		{
			name: "Success",
			path: "testdata/config.valid.yaml",
			wantData: &config.Configuration{
				DB: &config.Database{
					LogQueries: true,
					Timeout:    15,
				},
				Server: &config.Server{
					Port:         ":3000",
					Debug:        true,
					ReadTimeout:  10,
					WriteTimeout: 15,
				},
				App: &config.Application{
					SwaggerUIPath: "assets/swagger",
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := config.Load(tt.path)
			assert.Equal(t, tt.wantData, cfg)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
