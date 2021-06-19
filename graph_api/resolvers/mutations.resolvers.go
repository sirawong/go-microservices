package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/sirawong/bookrental-microservices/graph_api/generated"
	"github.com/sirawong/bookrental-microservices/graph_api/model"
	bookspb "github.com/sirawong/bookrental-microservices/proto/books"
	holderspb "github.com/sirawong/bookrental-microservices/proto/holders"
)

func (r *mutationResolver) CreateBook(ctx context.Context, inputData model.BookInput) (*model.Book, error) {
	log.Printf("Book: %+v", inputData)

	book := graph2ServiceBookInput(&inputData)

	addBookResponse, err := r.services.Books().AddBook(ctx, &bookspb.AddBookRequest{Book: book})
	if err != nil {
		return nil, err
	}

	return service2GraphBook(addBookResponse.Book), nil
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (bool, error) {
	_, err := r.services.Books().DeleteBook(ctx, &bookspb.DeleteBookRequest{Id: id})
	return err == nil, err
}

func (r *mutationResolver) TakeBookInUse(ctx context.Context, holderID string, bookID string) (bool, error) {
	book, holder, err := getBookAndHolder(ctx, r.services, holderID, bookID)
	if err != nil {
		return false, err
	}
	log.Printf("Found book: %+v", book)

	// Add book id to holder
	holder.HeldBooks = append(holder.HeldBooks, bookID)
	_, err = r.services.Holders().UpdateHolder(ctx, &holderspb.UpdateHolderRequest{Holder: holder})
	if err != nil {
		return false, err
	}
	log.Print("Holder updated successfully")

	return err == nil, err
}

func (r *mutationResolver) ReturnBook(ctx context.Context, holderID string, bookID string) (bool, error) {
	book, holder, err := getBookAndHolder(ctx, r.services, holderID, bookID)
	if err != nil {
		return false, err
	}
	log.Printf("Found book: %+v", book)

	for pos, id := range holder.HeldBooks {
		if bookID == id {
			holder.HeldBooks = append(holder.HeldBooks[:pos], holder.HeldBooks[pos+1:]...)
			break
		}
	}

	_, err = r.services.Holders().UpdateHolder(ctx, &holderspb.UpdateHolderRequest{Holder: holder})
	if err != nil {
		return false, err
	}
	log.Print("Holder updated successfully")

	return err == nil, err
}

func (r *mutationResolver) CreateHolder(ctx context.Context, inputData model.HolderInput) (*model.Holder, error) {
	log.Printf("Holder: %+v", inputData)

	holder := graph2ServiceHolderInput(&inputData)

	addHolderResponse, err := r.services.Holders().AddHolder(ctx, &holderspb.AddHolderRequest{Holder: holder})
	if err != nil {
		return nil, err
	}

	return service2GraphHolder(addHolderResponse.Holder), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
