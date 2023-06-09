package LoadEnv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		envName string
		want    string
	}{

		{
			envName: "API_PARK",
			want:    "PARK",
		},
		{
			envName: "API_END",
			want:    "end",
		},
		{
			envName: "API_BEGIN",
			want:    "BEGIN",
		},
		{
			envName: "API_FORMULA",
			want:    "FORMULA",
		},
		{
			envName: "API_KRAIMER",
			want:    "KRAIMER",
		},
		{
			envName: "API_REST",
			want:    "REST",
		},
		{
			envName: "API_FULL",
			want:    "FULL",
		},
		{
			envName: "API_GOLANG",
			want:    "GOLANG",
		},
		{
			envName: "API_JAVA",
			want:    "JAVA",
		},
		{
			envName: "API_KOTLIN",
			want:    "KOTLIN",
		},
		{
			envName: "API_MAIN",
			want:    "MAIN",
		},
		{
			envName: "API_JAVASCRIPT",
			want:    "JAVASCRIPT",
		},
		{
			envName: "API_RUBY",
			want:    "RUBY",
		},
		{
			envName: "API_PYTHON",
			want:    "PYTHON",
		},
		{
			envName: "API_HASKELL",
			want:    "HASKELL",
		},
		{
			envName: "API_LUA",
			want:    "LUA",
		},
		{
			envName: "API_PHP",
			want:    "PHP",
		},
		{
			envName: "API_ENV",
			want:    "ENV",
		},
		{
			envName: "API_C",
			want:    "C",
		},
		{
			envName: "API_COBOL",
			want:    "COBOL",
		},
		{
			envName: "API_ANDRE",
			want:    "ANDRE",
		},
		{
			envName: "API_MYSQL",
			want:    "MYSQL",
		},
		{
			envName: "API_PORT",
			want:    "PORT",
		},
		{
			envName: "API_BINDING_PORT",
			want:    "BINDING_PORT",
		},
		{
			envName: "API_DOCKER_BIND_PORT",
			want:    "DOCKER_BIND_PORT",
		}, // CUT
		{
			envName: "ENV_DOCKER",
			want:    "DOCKER",
		},
		{
			envName: "ENV_KUBERNETES",
			want:    "KUBERNETES",
		},
		{
			envName: "ENV_RUBY",
			want:    "RUBY",
		},
		{
			envName: "ENV_GO",
			want:    "GO",
		},
		{
			envName: "ENV_C",
			want:    "C",
		},
		{
			envName: "ENV_CSHARP",
			want:    "C#",
		},
		{
			envName: "ENV_PYTHON",
			want:    "PYTHON",
		},
		{
			envName: "ENV_ELIXIR",
			want:    "ELIXIR",
		},
		{
			envName: "ENV_PODS",
			want:    "PODS",
		},
		{
			envName: "ENV_HOST_NAME",
			want:    "HOST_NAME",
		},
		{
			envName: "ENV_CASSANDRA",
			want:    "CASSANDRA",
		},
		{
			envName: "ENV_SCYLLA",
			want:    "SCYLLA",
		},
		{
			envName: "ENV_VALHALLA",
			want:    "VALHALLA",
		},
		{
			envName: "ENV_PERL",
			want:    "PERL",
		},
		{
			envName: "ENV_DISCORD",
			want:    "DISCORD",
		},
		{
			envName: "ENV_PHP",
			want:    "PHP",
		},
		{
			envName: "ENV_LUA",
			want:    "LUA",
		},
		{
			envName: "ENV_ENV",
			want:    "ENV",
		},
		{
			envName: "ENV_COBOL",
			want:    "COBOL",
		},
		{
			envName: "ENV_PASCAL",
			want:    "PASCAL",
		},
		{
			envName: "ENV_RUST",
			want:    "RUST",
		},
		{
			envName: "ENV_MYSQL",
			want:    "MYSQL",
		},
		{
			envName: "ENV_PORT",
			want:    "PORT",
		},
		{
			envName: "ENV_NOSQL",
			want:    "NOSQL",
		},
		{
			envName: "ENV_MONGODB",
			want:    "MONGODB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.envName, func(t *testing.T) {
			got := LoadEnv(tt.envName)
			assert.Equal(t, tt.want, got)
		})
	}

}
