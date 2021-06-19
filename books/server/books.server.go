package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sirawong/bookrental-microservices/config"
	"github.com/sirawong/bookrental-microservices/db"
	"github.com/sirawong/bookrental-microservices/models"
	bookspb "github.com/sirawong/bookrental-microservices/proto/books"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	database        *mongo.Database
	booksCollection *mongo.Collection
}

func NewServer(ctx context.Context) *Server {

	configDB := config.GetDbConfig()
	client := db.ConnectDB(ctx, configDB)

	database := client.Database(configDB.DBName)
	collection := database.Collection(configDB.Collection)

	return &Server{
		database:        database,
		booksCollection: collection,
	}
}

func (s *Server) Run() {
	port := fmt.Sprintf(":%s", config.GetServiceConfig().AppPort)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Unable to listen on port %s | %v", port, err)
		return
	}
	grpcServer := grpc.NewServer()

	bookspb.RegisterBooksAPIServer(grpcServer, s)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("Failed to serve:", err)
		}
	}()
	log.Printf("Starting Books server on port %s", port)

}

func (s *Server) ListBooks(ctx context.Context, req *bookspb.ListBooksRequest) (*bookspb.ListBooksResponse, error) {
	// Initiate a BookItem type to write decoded data to
	data := models.BookModel{}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := s.booksCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())

	var booksList []*bookspb.Book
	// cursor.Next() returns a boolean, if false there are no more items and loop will break
	for cursor.Next(context.Background()) {
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		// If no error is found send Book over stream
		book := &bookspb.Book{
			Id:     data.ID.Hex(),
			Author: data.Author,
			Title:  data.Title,
		}

		booksList = append(booksList, book)

	}

	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown cursor error: %v", err))
	}

	return &bookspb.ListBooksResponse{Books: booksList}, nil
}

func (s *Server) GetBook(ctx context.Context, req *bookspb.GetBookRequest) (*bookspb.GetBookResponse, error) {
	// convert string id (from proto) to mongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := s.booksCollection.FindOne(ctx, bson.M{"_id": oid})
	// Create an empty BookItem to write our decode result to
	data := models.BookModel{}
	// decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find book with Object Id %s: %v", req.GetId(), err))
	}
	// Cast to ReadBookRes type
	response := &bookspb.GetBookResponse{
		Book: &bookspb.Book{
			Id:     oid.Hex(),
			Author: data.Author,
			Title:  data.Title,
		},
	}
	return response, nil
}

func (s *Server) AddBook(ctx context.Context, req *bookspb.AddBookRequest) (*bookspb.AddBookResponse, error) {
	// Essentially doing req.Book to access the struct with a nil check
	book := req.GetBook()
	// Now we have to convert this into a BookItem type to convert into BSON
	data := models.BookModel{
		// ID:    Empty, so it gets omitted and MongoDB generates a unique Object ID upon insertion.
		Author: book.GetAuthor(),
		Title:  book.GetTitle(),
	}

	// Insert the data into the database, result contains the newly generated Object ID for the new document
	result, err := s.booksCollection.InsertOne(ctx, data)
	// check for potential errors
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	// add the id to book, first cast the "generic type" (go doesn't have real generics yet) to an Object ID.
	oid := result.InsertedID.(primitive.ObjectID)
	// Convert the object id to it's string counterpart
	book.Id = oid.Hex()
	// return the book in a CreateBookRes type
	return &bookspb.AddBookResponse{Book: book}, nil
}

func (s *Server) GetBooks(ctx context.Context, req *bookspb.GetBooksRequest) (*bookspb.GetBooksResponse, error) {
	oids := req.GetIds()

	var booksList []*bookspb.Book

	for _, id := range oids {
		i, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
		}
		result := s.booksCollection.FindOne(ctx, bson.M{"_id": i})
		// Create an empty BookItem to write our decode result to
		data := models.BookModel{}
		// decode and write to data
		if err := result.Decode(&data); err != nil {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find book with Object Id %s: %v", i, err))
		}
		// Cast to ReadBookRes type
		book := &bookspb.Book{
			Id:     data.ID.Hex(),
			Author: data.Author,
			Title:  data.Title,
		}
		booksList = append(booksList, book)
	}

	return &bookspb.GetBooksResponse{Books: booksList}, nil

}

func (s *Server) DeleteBook(ctx context.Context, req *bookspb.DeleteBookRequest) (*bookspb.DeleteBookResponse, error) {
	// Get the ID (string) from the request message and convert it to an Object ID
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	// Check for errors
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	// DeleteOne returns DeleteResult which is a struct containing the amount of deleted docs (in this case only 1 always)
	// So we return a boolean instead
	_, err = s.booksCollection.DeleteOne(ctx, bson.M{"_id": oid})
	// Check for errors
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete book with id %s: %v", req.GetId(), err))
	}
	// Return response with success: true if no error is thrown (and thus document is removed)
	return &bookspb.DeleteBookResponse{}, nil
}
