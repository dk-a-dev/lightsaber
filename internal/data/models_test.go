package data

import (
	"testing"

	"lightsaber.dkadev.xyz/internal/validator"
)

func TestValidateMovie(t *testing.T) {
	tests := []struct {
		name    string
		movie   Movie
		wantErr bool
	}{
		{
			name: "valid movie",
			movie: Movie{
				Title:   "Test Movie",
				Year:    2023,
				Runtime: 120,
				Genres:  []string{"Action", "Drama"},
			},
			wantErr: false,
		},
		{
			name: "empty title",
			movie: Movie{
				Title:   "",
				Year:    2023,
				Runtime: 120,
				Genres:  []string{"Action"},
			},
			wantErr: true,
		},
		{
			name: "title too long",
			movie: Movie{
				Title:   string(make([]byte, 501)), // 501 characters
				Year:    2023,
				Runtime: 120,
				Genres:  []string{"Action"},
			},
			wantErr: true,
		},
		{
			name: "invalid year - too old",
			movie: Movie{
				Title:   "Test Movie",
				Year:    1800,
				Runtime: 120,
				Genres:  []string{"Action"},
			},
			wantErr: true,
		},
		{
			name: "invalid year - future",
			movie: Movie{
				Title:   "Test Movie",
				Year:    2030,
				Runtime: 120,
				Genres:  []string{"Action"},
			},
			wantErr: true,
		},
		{
			name: "invalid runtime - negative",
			movie: Movie{
				Title:   "Test Movie",
				Year:    2023,
				Runtime: -10,
				Genres:  []string{"Action"},
			},
			wantErr: true,
		},
		{
			name: "no genres",
			movie: Movie{
				Title:   "Test Movie",
				Year:    2023,
				Runtime: 120,
				Genres:  []string{},
			},
			wantErr: true,
		},
		{
			name: "too many genres",
			movie: Movie{
				Title:   "Test Movie",
				Year:    2023,
				Runtime: 120,
				Genres:  []string{"Action", "Drama", "Comedy", "Horror", "Romance", "Thriller"},
			},
			wantErr: true,
		},
		{
			name: "duplicate genres",
			movie: Movie{
				Title:   "Test Movie",
				Year:    2023,
				Runtime: 120,
				Genres:  []string{"Action", "Action"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.New()
			ValidateMovie(v, &tt.movie)

			if tt.wantErr && v.Valid() {
				t.Error("expected validation to fail")
			}
			if !tt.wantErr && !v.Valid() {
				t.Errorf("expected validation to pass, got errors: %v", v.Errors)
			}
		})
	}
}

func TestValidateUser(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name: "valid user with password set",
			user: func() User {
				u := User{
					Name:  "John Doe",
					Email: "john@example.com",
				}
				u.Password.Set("validpassword123")
				return u
			}(),
			wantErr: false,
		},
		{
			name: "empty name",
			user: func() User {
				u := User{
					Name:  "",
					Email: "john@example.com",
				}
				u.Password.Set("validpassword123")
				return u
			}(),
			wantErr: true,
		},
		{
			name: "name too long",
			user: func() User {
				u := User{
					Name:  string(make([]byte, 501)),
					Email: "john@example.com",
				}
				u.Password.Set("validpassword123")
				return u
			}(),
			wantErr: true,
		},
		{
			name: "invalid email",
			user: func() User {
				u := User{
					Name:  "John Doe",
					Email: "invalid-email",
				}
				u.Password.Set("validpassword123")
				return u
			}(),
			wantErr: true,
		},
		{
			name: "empty email",
			user: func() User {
				u := User{
					Name:  "John Doe",
					Email: "",
				}
				u.Password.Set("validpassword123")
				return u
			}(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.New()
			ValidateUser(v, &tt.user)

			if tt.wantErr && v.Valid() {
				t.Error("expected validation to fail")
			}
			if !tt.wantErr && !v.Valid() {
				t.Errorf("expected validation to pass, got errors: %v", v.Errors)
			}
		})
	}
}
