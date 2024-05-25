package user

import (
	"context"
	"gophkeeper/internal/server/model"
)

// FindByEmail retrieves a user from the database based on their email.
// It takes a context and an email as parameters and returns the found User and an error, if any.
// The User struct contains fields for ID, Email, and Password.
// The function executes a SQL query to select the user with the given email,
// scans the result into the User struct, and returns it along with any error encountered.
func (r *Repository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := r.db.QueryRowContext(
		ctx,
		`SELECT id, email, password FROM users WHERE email = $1`,
		email,
	).Scan(&user.ID, &user.Email, &user.Password)

	return user, err
}
