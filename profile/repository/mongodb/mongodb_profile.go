package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"screener.com/domain"
)

const (
	profilesCollection = "profiles"
	profilerDatabase   = "profiler"
)

type profileRepository struct {
	Client *mongo.Client
}

// NewProfileRepository returns a fresh profileRepository
func NewProfileRepository(client *mongo.Client) domain.ProfileRepository {
	return &profileRepository{client}
}

// GetFullCIKList returns the full list of available CIKs in profiles collection in db
func (r *profileRepository) GetFullCIKList(ctx context.Context) (*[]interface{}, error) {
	// TODO: to write unit tests, need to find a way to mock mongo.Client
	// and mongo.Collection
	collection := r.Client.Database(profilerDatabase).Collection(profilesCollection)
	opts := options.Distinct()
	filter := bson.D{}

	ciks, err := collection.Distinct(ctx, "cik", filter, opts)
	if err != nil {
		return nil, err
	}

	// TODO: is there a way in which the user does not have to worry about coercing
	// this array's values into strings?
	return &ciks, nil
}

// GetFullProfileForCIK returns the full financial profile for proviede CIK
// spanning available range of years
func (r *profileRepository) GetFullProfileForCIK(ctx context.Context, cik string) (*domain.FullCompanyProfile, error) {
	collection := r.Client.Database(profilerDatabase).Collection(profilesCollection)
	filter := bson.D{{"cik", cik}}
	opts := options.Find()
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var results domain.FullCompanyProfile
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return &results, nil
}

// GetFullProfileForTicker returns the full financial profile for provided ticker
// spanning available range of years
func (r *profileRepository) GetFullProfileForTicker(ctx context.Context, ticker string) (*domain.FullCompanyProfile, error) {
	collection := r.Client.Database(profilerDatabase).Collection(profilesCollection)
	filter := bson.D{{"ticker", ticker}}
	opts := options.Find()
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var results domain.FullCompanyProfile
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return &results, nil
}
