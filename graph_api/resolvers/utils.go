package resolvers

import (
	"context"
	"log"

	"github.com/sirawong/bookrental-microservices/graph_api/model"
	"github.com/sirawong/bookrental-microservices/graph_api/services"
	bookspb "github.com/sirawong/bookrental-microservices/proto/books"
	holderspb "github.com/sirawong/bookrental-microservices/proto/holders"
)

func service2GraphBook(book *bookspb.Book) *model.Book {
	return &model.Book{
		ID:     book.Id,
		Author: book.Author,
		Title:  book.Title,
	}
}

func graph2ServiceHolderInput(holderInput *model.HolderInput) *holderspb.Holder {
	return &holderspb.Holder{
		FirstName: softDeference(holderInput.FirstName),
		LastName:  softDeference(holderInput.LastName),
		Email:     softDeference(holderInput.Email),
	}
}

func graph2ServiceBookInput(bookInput *model.BookInput) *bookspb.Book {
	return &bookspb.Book{
		Id:     softDeference(bookInput.ID),
		Author: softDeference(bookInput.Author),
		Title:  softDeference(bookInput.Title),
	}
}
func service2GraphHolder(holder *holderspb.Holder) *model.Holder {
	return &model.Holder{
		ID:        holder.Id,
		FirstName: holder.FirstName,
		LastName:  holder.LastName,
		Email:     holder.Email,
		HeldBooks: holder.HeldBooks,
	}
}

func softDeference(field *string) string {
	if field == nil {
		return ""
	}
	return *field
}

func getBookAndHolder(ctx context.Context, s services.Services, holderId, bookId string) (*bookspb.Book, *holderspb.Holder, error) {
	log.Printf("Request data. HolderID: %s, bookID: %s", holderId, bookId)

	getHolderResponse, err := s.Holders().GetHolder(ctx, &holderspb.GetHolderRequest{Id: holderId})
	if err != nil {
		return nil, nil, err
	}
	getBookResponse, err := s.Books().GetBook(ctx, &bookspb.GetBookRequest{Id: bookId})
	if err != nil {
		return nil, nil, err
	}

	return getBookResponse.Book, getHolderResponse.Holder, nil
}
