package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sirawong/bookrental-microservices/config"
	"github.com/sirawong/bookrental-microservices/db"
	"github.com/sirawong/bookrental-microservices/models"
	holderspb "github.com/sirawong/bookrental-microservices/proto/holders"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	database          *mongo.Database
	holdersCollection *mongo.Collection
}

func NewServer(ctx context.Context) *Server {

	configDB := config.GetDbConfig()
	client := db.ConnectDB(ctx, configDB)

	database := client.Database(configDB.DBName)
	collection := database.Collection(configDB.Collection)

	return &Server{
		database:          database,
		holdersCollection: collection,
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

	holderspb.RegisterHoldersAPIServer(grpcServer, s)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("Failed to serve:", err)
		}
	}()
	log.Printf("Starting Books server on port %s", port)
}

// Get all customers
func (s *Server) ListHolders(ctx context.Context, req *holderspb.ListHoldersRequest) (*holderspb.ListHoldersResponse, error) {
	// Initiate a BookItem type to write decoded data to
	data := models.HolderModel{}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := s.holdersCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())

	var holdersList []*holderspb.Holder
	// cursor.Next() returns a boolean, if false there are no more items and loop will break
	for cursor.Next(ctx) {
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		// If no error is found send Book over stream
		holder := &holderspb.Holder{
			Id:        data.ID.Hex(),
			FirstName: data.First_name,
			LastName:  data.Last_name,
			Email:     data.Email,
			HeldBooks: data.Held_books,
		}

		holdersList = append(holdersList, holder)

	}

	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown cursor error: %v", err))
	}

	return &holderspb.ListHoldersResponse{Holders: holdersList}, nil
}

// Get customer by book ID
func (s *Server) GetHolderByBookId(ctx context.Context, req *holderspb.GetHolderByBookIdRequest) (*holderspb.GetHolderByBookIdResponse, error) {
	// convert string id (from proto) to mongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	data := models.HolderModel{}
	result := s.holdersCollection.FindOne(ctx, bson.D{{"held_books", bson.A{oid}}})
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find book with Object Id %s: %v", req.GetId(), err))
	}
	// Cast to ReadBookRes type
	return &holderspb.GetHolderByBookIdResponse{
		Holder: &holderspb.Holder{
			Id:        oid.Hex(),
			FirstName: data.First_name,
			LastName:  data.Last_name,
			Email:     data.Email,
			HeldBooks: data.Held_books,
		},
	}, nil

}

// Get customer by ID
func (s *Server) GetHolder(ctx context.Context, req *holderspb.GetHolderRequest) (*holderspb.GetHolderResponse, error) {
	// convert string id (from proto) to mongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := s.holdersCollection.FindOne(ctx, bson.M{"_id": oid})
	// Create an empty BookItem to write our decode result to
	data := models.HolderModel{}
	// decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find book with Object Id %s: %v", req.GetId(), err))
	}
	// Cast to ReadBookRes type
	response := &holderspb.GetHolderResponse{
		Holder: &holderspb.Holder{
			Id:        oid.Hex(),
			FirstName: data.First_name,
			LastName:  data.Last_name,
			Email:     data.Email,
			HeldBooks: data.Held_books,
		},
	}
	return response, nil
}

// Add new customer
func (s *Server) AddHolder(ctx context.Context, req *holderspb.AddHolderRequest) (*holderspb.AddHolderResponse, error) {
	// Essentially doing req.Book to access the struct with a nil check
	holder := req.GetHolder()
	// Now we have to convert this into a BookItem type to convert into BSON
	data := models.HolderModel{
		// ID:    Empty, so it gets omitted and MongoDB generates a unique Object ID upon insertion.
		First_name: holder.GetFirstName(),
		Last_name:  holder.GetLastName(),
		Email:      holder.GetEmail(),
		Held_books: holder.GetHeldBooks(),
	}

	// Insert the data into the database, result contains the newly generated Object ID for the new document
	result, err := s.holdersCollection.InsertOne(ctx, data)
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
	holder.Id = oid.Hex()
	// return the book in a CreateBookRes type
	return &holderspb.AddHolderResponse{Holder: holder}, nil
}

// Update customer
func (s *Server) UpdateHolder(ctx context.Context, req *holderspb.UpdateHolderRequest) (*holderspb.UpdateHolderResponse, error) {
	// Get the blog data from the request
	holder := req.GetHolder()

	// Convert the Id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(holder.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied blog id to a MongoDB ObjectId: %v", err),
		)
	}

	// Convert the oid into an unordered bson document to search by id
	filter := bson.M{"_id": oid}

	cursor, err := s.holdersCollection.Find(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	decoded := models.HolderModel{}
	err = cursor.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find blog with supplied ID: %v", err),
		)
	}

	heldBooks := decoded.Held_books
	heldBooks = append(heldBooks, holder.GetHeldBooks()...)

	// Convert the data to be updated into an unordered Bson document
	update := bson.M{
		"first_name": holder.GetFirstName(),
		"last_name":  holder.GetLastName(),
		"email":      holder.GetEmail(),
		"held_books": heldBooks,
	}

	// Result is the BSON encoded result
	// To return the updated document instead of original we have to add options.
	result, err := s.holdersCollection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Update error: %v", err))
	}
	log.Println(result.ModifiedCount)
	return &holderspb.UpdateHolderResponse{
		Holder: &holderspb.Holder{
			Id:        decoded.ID.Hex(),
			FirstName: holder.GetFirstName(),
			LastName:  holder.GetLastName(),
			Email:     holder.GetEmail(),
			HeldBooks: heldBooks,
		},
	}, nil
}
