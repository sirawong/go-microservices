package services

import (
	"log"

	bookspb "github.com/sirawong/bookrental-microservices/proto/books"
	holderspb "github.com/sirawong/bookrental-microservices/proto/holders"
	"google.golang.org/grpc"
)

type ServicesConfig struct {
	BooksSvc   string
	HoldersSvc string
}

type services struct {
	// io.Closer
	booksClientConn   *grpc.ClientConn
	booksClient       bookspb.BooksAPIClient
	holdersClientConn *grpc.ClientConn
	holdersClient     holderspb.HoldersAPIClient
}

type Services interface {
	Books() bookspb.BooksAPIClient
	Holders() holderspb.HoldersAPIClient
}

func NewClientGrpc(conf ServicesConfig) (Services, error) {
	log.Printf("Connection to Books Services: %s...", conf.BooksSvc)
	booksConnection, err := grpc.Dial(conf.BooksSvc, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	log.Printf("Connection to Golders Serice: %s...", conf.HoldersSvc)
	holdersConnection, err := grpc.Dial(conf.HoldersSvc, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	ah := &services{
		booksClientConn:   booksConnection,
		booksClient:       bookspb.NewBooksAPIClient(booksConnection),
		holdersClientConn: holdersConnection,
		holdersClient:     holderspb.NewHoldersAPIClient(holdersConnection),
	}
	return ah, nil
}

func (ah *services) Books() bookspb.BooksAPIClient {
	return ah.booksClient
}

func (ah *services) Holders() holderspb.HoldersAPIClient {
	return ah.holdersClient
}

func (ah *services) Close() error {
	err := ah.booksClientConn.Close()
	if err != nil {
		log.Println("An error occurred while closing connection on Books services:", err)
	}

	err = ah.holdersClientConn.Close()
	if err != nil {
		log.Println("An error occurred while closing connection on Holders services:", err)
	}
	return nil
}
