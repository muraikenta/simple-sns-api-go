package post

import (
	"context"
	"simple_sns_api/src/db"
	"simple_sns_api/src/ent"
	"simple_sns_api/src/ent/post"
)

type PostService struct{}

type PaginationParams struct {
	Cursor int
	Size   int
}

type CreateParams struct {
	UserId int
	Body   string
}

func (s PostService) find(ctx context.Context, pagination PaginationParams) ([]*ent.Post, error) {
	query := db.Client.Post.Query()
	if pagination.Cursor != 0 {
		query = query.Where(post.IDLT(pagination.Cursor))
	}
	if pagination.Size == 0 {
		pagination.Size = 20
	}
	posts, err := query.
		Limit(pagination.Size).
		Order(ent.Desc(post.FieldID)).
		All(ctx)
	return posts, err
}

func (s PostService) Create(ctx context.Context, params CreateParams) (*ent.Post, error) {
	post, err := db.Client.Post.Create().
		SetBody(params.Body).
		SetUserID(params.UserId).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return post, nil
}
