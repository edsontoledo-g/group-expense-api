package auth

import (
	"context"
	"time"

	"github.com/edsontoledo-g/group-expense-api/internal/domain/users"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository interface {
	CreateUser(user *users.User, auth *AuthProvider) error
	UpdateUser(user *users.User, auth *AuthProvider) error
	GetAuthByEmail(email string) (*AuthProvider, error)
	GetUserByVerificationToken(tokenHash string) (*users.User, *AuthProvider, error)
}

type authRepository struct {
	pool *pgxpool.Pool
}

func (repo *authRepository) CreateUser(user *users.User, auth *AuthProvider) error {
	ctx := context.Background()
	tx, err := repo.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	gs := pgx.NamedArgs{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"imageURL":  user.ImageURL,
	}
	query := `
		INSERT INTO users (first_name, last_name, email, image_url)
		VALUES (@firstName, @lastName, @email, @imageURL)
		RETURNING id, created_at, updated_at, is_verified
	`
	err = tx.QueryRow(ctx, query, gs).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.IsVerified)
	if err != nil {
		return err
	}
	auth.UserID = user.ID
	err = repo.createAuthProvider(auth, ctx, tx)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (repo *authRepository) UpdateUser(user *users.User, auth *AuthProvider) error {
	ctx := context.Background()
	tx, err := repo.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	gs := pgx.NamedArgs{
		"firstName":  user.FirstName,
		"lastName":   user.LastName,
		"email":      user.Email,
		"imageURL":   user.ImageURL,
		"updatedAt":  time.Now(),
		"isVerified": user.IsVerified,
	}
	query := `
		UPDATE users
		SET first_name = @firstName,
			last_name = @lastName,
			email = @email,
			image_url = @imageURL,
			updated_at = @updatedAt,
			is_verified = @isVerified
		WHERE user_id = @userID
	`
	_, err = tx.Exec(ctx, query, gs)
	if err != nil {
		return err
	}
	err = repo.updateAuthProvider(auth, ctx, tx)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (repo *authRepository) GetAuthByEmail(email string) (*AuthProvider, error) {
	ctx := context.Background()
	gs := pgx.NamedArgs{
		"email": email,
	}
	query := `SELECT id FROM users where email = @email`
	var userID uint
	err := repo.pool.QueryRow(ctx, query, gs).Scan(&userID)
	if err != nil {
		return nil, err
	}
	authProvider, err := repo.getAuthByUserID(userID, ctx)
	if err != nil {
		return nil, err
	}
	return authProvider, nil
}

func (repo *authRepository) GetUserByVerificationToken(tokenHash string) (*users.User, *AuthProvider, error) {
	ctx := context.Background()
	gs := pgx.NamedArgs{
		"verificationTokenHash": tokenHash,
		"currentTime":           time.Now(),
	}
	query := `
		SELECT user_id
		FROM auth_providers
		WHERE verification_token_hash = @verificationTokenHash AND verification_token_expires_at > @currentTime
		LIMIT 1
	`
	auth := &AuthProvider{}
	err := repo.pool.QueryRow(ctx, query, gs).Scan(
		&auth.ID,
		&auth.UserID,
		&auth.Provider,
		&auth.ProviderUserID,
		&auth.PasswordHash,
		&auth.VerificationTokenHash,
		&auth.VerificationTokenExpiresAt,
		&auth.CreatedAt,
	)
	if err != nil {
		return nil, nil, err
	}
	user, err := repo.getUserByUserID(auth.UserID)
	if err != nil {
		return nil, nil, err
	}
	return user, auth, nil
}

func (repo *authRepository) createAuthProvider(auth *AuthProvider, ctx context.Context, tx pgx.Tx) error {
	gs := pgx.NamedArgs{
		"userID":                     auth.UserID,
		"provider":                   auth.Provider,
		"providerUserID":             auth.ProviderUserID,
		"passwordHash":               auth.PasswordHash,
		"verificationTokenHash":      auth.VerificationTokenHash,
		"verificationTokenExpiresAt": auth.VerificationTokenExpiresAt,
	}
	query := `
		INSERT INTO auth_providers (user_id, provider, provider_user_id, password_hash, verification_token_hash, verification_token_expires_at)
		VALUES (@userID, @provider, @providerUserID, @passwordHash, @verificationTokenHash, @verificationTokenExpiresAt)
	`
	_, err := tx.Exec(ctx, query, gs)
	return err
}

func (repo *authRepository) updateAuthProvider(auth *AuthProvider, ctx context.Context, tx pgx.Tx) error {
	gs := pgx.NamedArgs{
		"userID":                     auth.UserID,
		"provider":                   auth.Provider,
		"providerUserID":             auth.ProviderUserID,
		"passwordHash":               auth.PasswordHash,
		"verificationTokenHash":      auth.VerificationTokenHash,
		"verificationTokenExpiresAt": auth.VerificationTokenExpiresAt,
	}
	query := `
		UPDATE auth_providers
		SET provider = @provider,
			provider_user_id = @providerUserID,
			password_hash = @passwordHash,
			verification_token_hash = @verificationTokenHash,
			verification_token_expires_at = @verificationTokenExpiresAt
		WHERE user_id = @userID
	`
	_, err := tx.Exec(ctx, query, gs)
	return err
}

func (repo *authRepository) getAuthByUserID(userID uint, ctx context.Context) (*AuthProvider, error) {
	gs := pgx.NamedArgs{
		"userID": userID,
	}
	query := `
		SELECT id, user_id, provider, provider_user_id, password_hash, verification_token_hash, verification_token_expires_at, created_at
		FROM auth_providers
		WHERE user_id = @userID
	`
	authProvider := &AuthProvider{}
	err := repo.pool.QueryRow(ctx, query, gs).Scan(
		&authProvider.ID,
		&authProvider.UserID,
		&authProvider.Provider,
		&authProvider.ProviderUserID,
		&authProvider.PasswordHash,
		&authProvider.VerificationTokenHash,
		&authProvider.VerificationTokenExpiresAt,
		&authProvider.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return authProvider, nil
}

func (repo *authRepository) getUserByUserID(id uint) (*users.User, error) {
	ctx := context.Background()
	gs := pgx.NamedArgs{
		"userID": id,
	}
	query := `
		SELECT id, first_name, last_name, email, image_url, created_at, updated_at, is_verified
		FROM users
		WHERE id = @userID
	`
	user := &users.User{}
	err := repo.pool.QueryRow(ctx, query, gs).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.ImageURL,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsVerified,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return &authRepository{
		pool: db,
	}
}
