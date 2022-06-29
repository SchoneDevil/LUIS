package users

import (
	"testing"

	repo "user_service/internal/usecase/users/repo"
)

func TestLogin(t *testing.T) {
	type args struct {
		email, password string
	}
	repo := repo.NewMockDB()
	srv := LoadUserService(repo)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful login",
			args: args{
				email:    "test@example.com",
				password: "content",
			},
			wantErr: false,
		},
		{
			name: "Failed password incorrect",
			args: args{
				email:    "test@example.com",
				password: "fail_password",
			},
			wantErr: true,
		},
		{
			name: "Failed email not found",
			args: args{
				email:    "fake@example.com",
				password: "content",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run("test login with "+tt.name, func(t *testing.T) {
			_, err := srv.Login(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("error while executing login operation. error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestGetUserById(t *testing.T) {
	type args struct {
		id string
	}
	repo := repo.NewMockDB()
	srv := LoadUserService(repo)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Successful get user",
			args:    args{id: "be811bfb-ea85-405e-975f-df40dae20bad"},
			wantErr: false,
		},
		{
			name:    "Failed get user",
			args:    args{id: "37ef2e0c-2320-4bbd-91c0-647957bda0b1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run("test get user by id with "+tt.name, func(t *testing.T) {
			_, err := srv.GetUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("error while executing get user by id operation. error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestGetUserByEmail(t *testing.T) {
	type args struct {
		email string
	}
	repo := repo.NewMockDB()
	srv := LoadUserService(repo)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Successful get user",
			args:    args{email: "test@example.com"},
			wantErr: false,
		},
		{
			name:    "Failed get user",
			args:    args{email: "fail@example.com"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run("test get user by email with "+tt.name, func(t *testing.T) {
			_, err := srv.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("error while executing get user by email operation. error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestGetUsers(t *testing.T) {
	type args struct {
		id string
	}
	repo := repo.NewMockDB()
	srv := LoadUserService(repo)

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Successful get users",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run("test get users with "+tt.name, func(t *testing.T) {
			users := srv.GetUsers()
			if (len(users) == 0) != tt.wantErr {
				t.Errorf("error while executing get users operation. error = %v, wantErr %v", "err", tt.wantErr)
			}
		})
	}
}
func TestAddUser(t *testing.T) {
	type args struct {
		dto CreateUserDTO
	}
	repo := repo.NewMockDB()
	srv := LoadUserService(repo)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful create user",
			args: args{dto: CreateUserDTO{
				Username: "new user",
				Email:    "user@example.com",
				Password: "123456",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run("test add user with "+tt.name, func(t *testing.T) {
			err := srv.CreateUser(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("error while executing add user operation. error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestUpdateUser(t *testing.T) {
	type args struct {
		id  string
		dto UpdateUserDTO
	}
	repo := repo.NewMockDB()
	srv := LoadUserService(repo)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful update user",
			args: args{
				id: "be811bfb-ea85-405e-975f-df40dae20bad",
				dto: UpdateUserDTO{
					Username: "new user",
					Email:    "user@example.com",
				},
			},
			wantErr: false,
		},
		{
			name: "Failed update user",
			args: args{
				id: "37ef2e0c-2320-4bbd-91c0-647957bda0b1",
				dto: UpdateUserDTO{
					Username: "new user",
					Email:    "user@example.com",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run("test update user with "+tt.name, func(t *testing.T) {
			err := srv.UpdateUser(tt.args.id, tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("error while executing update user operation. error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestDeleteUser(t *testing.T) {
	type args struct {
		id string
	}
	repo := repo.NewMockDB()
	srv := LoadUserService(repo)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful delete user",
			args: args{
				id: "be811bfb-ea85-405e-975f-df40dae20bad",
			},
			wantErr: false,
		},
		{
			name: "Failed delete user",
			args: args{
				id: "37ef2e0c-2320-4bbd-91c0-647957bda0b2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run("test delete user with "+tt.name, func(t *testing.T) {
			err := srv.DeleteUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("error while executing delete user operation. error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
