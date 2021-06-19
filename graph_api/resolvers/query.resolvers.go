package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/sirawong/bookrental-microservices/graph_api/generated"
	"github.com/sirawong/bookrental-microservices/graph_api/model"
	bookspb "github.com/sirawong/bookrental-microservices/proto/books"
	holderspb "github.com/sirawong/bookrental-microservices/proto/holders"
)

type queryResolver struct{ *Resolver }

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *queryResolver) Books(ctx context.Context, id *string, holderID *string) ([]*model.Book, error) {
	books := []*model.Book{}

	if id != nil {
		log.Printf("Book id: %s", *id)
		getBookResponse, err := r.services.Books().GetBook(ctx, &bookspb.GetBookRequest{Id: *id})
		if err != nil {
			return nil, err
		}
		books = append(books, service2GraphBook(getBookResponse.Book))
	} else if holderID != nil {
		getHolderResponse, err := r.services.Holders().GetHolder(ctx, &holderspb.GetHolderRequest{Id: *holderID})
		if err != nil {
			return nil, err
		}

		for _, bookId := range getHolderResponse.Holder.HeldBooks {
			getBookResponse, err := r.services.Books().GetBook(ctx, &bookspb.GetBookRequest{Id: bookId})
			if err != nil {
				return nil, err
			}
			books = append(books, service2GraphBook(getBookResponse.Book))
		}

	} else {
		listBookResponse, err := r.services.Books().ListBooks(ctx, &bookspb.ListBooksRequest{})
		if err != nil {
			return nil, err
		}
		for _, book := range listBookResponse.Books {
			books = append(books, service2GraphBook(book))
		}
		fmt.Println(books)
	}

	return books, nil
}

func (r *queryResolver) Holders(ctx context.Context, id *string, bookID *string) ([]*model.Holder, error) {
	holders := []*model.Holder{}

	if id != nil {
		log.Printf("Holder id: %s", *id)
		getHolderResponse, err := r.services.Holders().GetHolder(ctx, &holderspb.GetHolderRequest{Id: *id})
		if err != nil {
			return nil, err
		}
		holders = append(holders, service2GraphHolder(getHolderResponse.Holder))
	} else if bookID != nil {
		log.Printf("Book id: %s", *bookID)
		getHolderByBookIdResponse, err := r.services.Holders().GetHolderByBookId(ctx, &holderspb.GetHolderByBookIdRequest{Id: *bookID})
		if err != nil {
			return nil, err
		}
		holders = append(holders, service2GraphHolder(getHolderByBookIdResponse.Holder))
	} else {
		log.Printf("ALl books will be retrieved")
		listHoldersResponse, err := r.services.Holders().ListHolders(ctx, &holderspb.ListHoldersRequest{})
		if err != nil {
			return nil, err
		}
		for _, holder := range listHoldersResponse.Holders {
			holders = append(holders, service2GraphHolder(holder))
		}
	}

	return holders, nil
}
