package repository

import (
	"context"
	"go-starter/internal/data"
	"go-starter/internal/data/ent"
	"go-starter/internal/data/ent/sample"
	"go-starter/internal/data/structs"
	"go-starter/pkg/cache"
	"go-starter/pkg/log"
	"go-starter/pkg/validator"

	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"
)

// Sample represents the sample repository interface.
type Sample interface {
	Hello(ctx context.Context, body structs.Sample) (*ent.Sample, error)
}

// sampleRepo implements the Sample interface.
type sampleRepo struct {
	ec *ent.Client
	rc *redis.Client
	ms *meilisearch.Client
	c  *cache.Cache[ent.Sample]
}

// NewSample creates a new sample repository.
func NewSample(d *data.Data) Sample {
	entClient := d.GetEntClient()
	redisClient := d.GetRedis()
	meiliClient := d.GetMeilisearch()
	cacheInstance := cache.NewCache[ent.Sample](redisClient, cache.Key("sc_sample"), false)
	return &sampleRepo{ec: entClient, rc: redisClient, ms: meiliClient, c: cacheInstance}
}

func (r *sampleRepo) Hello(ctx context.Context, p structs.Sample) (*ent.Sample, error) {
	// try to fetch from cache.
	cf := p.Name
	row, err := r.c.Get(ctx, cf)
	if validator.IsNotNil(err) {
		// fetch from db when cache is empty.
		// use internal get method.
		row, err = r.getSample(ctx, &structs.FindSample{
			ID:   p.ID,
			Name: p.Name,
		})
		if validator.IsNotNil(err) {
			return nil, err
		}
		// set the cache.
		err := r.c.Set(ctx, cf, row)
		if err != nil {
			log.Errorf(nil, "failed to set cache: %v", err)
		}
	}
	return row, err
}

// CountX - sample count.
func (r *sampleRepo) CountX(ctx context.Context, p *structs.ListSamples) int {
	// create list builder
	builder, err := r.listBuilder(ctx, p)
	if validator.IsNotNil(err) {
		return 0
	}
	return builder.CountX(ctx)
}

// listBuilder - create list builder.
// internal method.
func (r *sampleRepo) listBuilder(ctx context.Context, p *structs.ListSamples) (*ent.SampleQuery, error) {
	// verify query params.
	var nextSample *ent.Sample
	if validator.IsNotEmpty(p.Cursor) {
		// query the sample.
		// use internal get method.
		row, err := r.getSample(ctx, &structs.FindSample{
			ID: p.Cursor,
		})
		if validator.IsNotNil(err) || validator.IsNil(row) {
			return nil, err
		}
		nextSample = row
	}

	// create builder.
	builder := r.ec.Sample.Query()

	// lt the cursor create time
	if nextSample != nil {
		builder.Where(sample.CreatedAtLT(nextSample.CreatedAt))
	}

	// set where conditions.
	// if validator.IsNotEmpty(p.User) {
	// 	builder.Where(sample.CreatedByEQ(p.User))
	// }
	// if validator.IsTrue(p.Private) || validator.IsFalse(p.Private) {
	// 	builder.Where(sample.PrivateEQ(p.Private))
	// }
	// if validator.IsTrue(p.Temp) || validator.IsFalse(p.Temp) {
	// 	builder.Where(sample.TempEQ(p.Temp))
	// }

	return builder, nil
}

// IsExist - check if sample exists
func (r *sampleRepo) IsExist(ctx context.Context, p *structs.FindSample) bool {
	// query the sample.
	// use internal get method.
	_, err := r.getSample(ctx, p)
	return validator.IsNotNil(err)
}

// getSample - get sample.
// internal method.
func (r *sampleRepo) getSample(ctx context.Context, p *structs.FindSample) (*ent.Sample, error) {
	// create builder.
	builder := r.ec.Sample.Query()

	// set where conditions.
	if validator.IsNotEmpty(p.ID) {
		builder.Where(sample.IDEQ(p.ID))
	}
	if validator.IsNotEmpty(p.Name) {
		builder.Where(sample.NameEQ(p.Name))
	}

	// execute the builder.
	row, err := builder.First(ctx)
	if validator.IsNotNil(err) {
		return nil, err
	}
	return row, nil
}

// CreateIndex creates a new index in Meilisearch.
func (r *sampleRepo) CreateIndex(ctx context.Context, sample *ent.Sample) error {
	// Get Meilisearch client
	ms := r.ms

	// Create Meilisearch index
	index := ms.Index("samples")

	// Define the document to index
	doc := map[string]any{
		"id":      sample.ID,
		"name":    sample.Name,
		"content": sample.Content,
		// Add other fields as needed
	}

	// Index the document
	_, err := index.AddDocuments([]map[string]interface{}{doc})
	if err != nil {
		log.Errorf(nil, "sampleRepo.CreateIndex error: %v\n", err)
		return err
	}

	return nil
}
