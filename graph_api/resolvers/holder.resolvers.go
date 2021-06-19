package resolvers

import (
	"context"
	"time"

	"github.com/sirawong/bookrental-microservices/graph_api/generated"
	"github.com/sirawong/bookrental-microservices/graph_api/model"
	bookspb "github.com/sirawong/bookrental-microservices/proto/books"
)

type holderResolver struct{ *Resolver }

// Holder returns gen.HolderResolver implementation.
func (r *Resolver) Holder() generated.HolderResolver { return &holderResolver{r} }

func (r *holderResolver) HeldBooks(ctx context.Context, obj *model.Holder) ([]*model.Book, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFn()

	resp, err := r.services.Books().GetBooks(ctx, &bookspb.GetBooksRequest{Ids: obj.HeldBooks})
	if err != nil {
		return nil, err
	}

	books := make([]*model.Book, len(resp.Books))
	for i, book := range resp.Books {
		books[i] = service2GraphBook(book)
	}

	return books, nil
}
