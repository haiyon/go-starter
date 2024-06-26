// Code generated by ent, DO NOT EDIT.

package ent

import (
	"go-starter/internal/data/ent/sample"
	"go-starter/internal/data/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	sampleMixin := schema.Sample{}.Mixin()
	sampleMixinFields0 := sampleMixin[0].Fields()
	_ = sampleMixinFields0
	sampleMixinFields3 := sampleMixin[3].Fields()
	_ = sampleMixinFields3
	sampleFields := schema.Sample{}.Fields()
	_ = sampleFields
	// sampleDescCreatedAt is the schema descriptor for created_at field.
	sampleDescCreatedAt := sampleMixinFields3[0].Descriptor()
	// sample.DefaultCreatedAt holds the default value on creation for the created_at field.
	sample.DefaultCreatedAt = sampleDescCreatedAt.Default.(func() time.Time)
	// sampleDescUpdatedAt is the schema descriptor for updated_at field.
	sampleDescUpdatedAt := sampleMixinFields3[1].Descriptor()
	// sample.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sample.DefaultUpdatedAt = sampleDescUpdatedAt.Default.(func() time.Time)
	// sample.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sample.UpdateDefaultUpdatedAt = sampleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sampleDescID is the schema descriptor for id field.
	sampleDescID := sampleMixinFields0[0].Descriptor()
	// sample.DefaultID holds the default value on creation for the id field.
	sample.DefaultID = sampleDescID.Default.(func() string)
}
