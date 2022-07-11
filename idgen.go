// Package idgen provides basic utilities for generating random identifiers.
package idgen

import (
	"fmt"
	"math/rand"
)

// These make sure that randomly generated IDs have enough entropy in them.
const (
	minCharSetSize = 10
	minLength      = 10
)

type options struct {
	charSet       []rune
	defaultLength int
}

// Generator generates random identifiers from a given character set using a
// pre-specified source of randomness.
type Generator struct {
	charSet       []rune
	defaultLength int
	r             *rand.Rand
}

type Option func(*options)

// WithCharSet specifies the set of characters to use when generating random
// identifiers. If not provided, identifiers default to case-sensitive alphanumerics.
func WithCharSet(charSet []rune) Option {
	return func(opts *options) {
		opts.charSet = charSet
	}
}

// WithDefaultLength sets the default length of the random identifiers to
// generate. If not provided, identifiers will be 20 characters long by default.
func WithDefaultLength(l int) Option {
	return func(opts *options) {
		opts.defaultLength = l
	}
}

// New initializes a new ID generator for the given source of randomness. For
// security-sensitive applications, be sure to use a source of randomness
// backed by a cryptographically secure random number generator, like
// github.com/Silicon-Ally/cryptorand.
func New(r *rand.Rand, opts ...Option) (*Generator, error) {
	randOpts := &options{
		charSet:       []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"),
		defaultLength: 20,
	}
	for _, opt := range opts {
		opt(randOpts)
	}

	if len(randOpts.charSet) < minCharSetSize {
		return nil, fmt.Errorf("character set is too small, needs at least %d characters", minCharSetSize)
	}

	if randOpts.defaultLength < minLength {
		return nil, fmt.Errorf("default ID length is too small, needs to be at least %d characters long", minLength)
	}

	return &Generator{
		charSet:       randOpts.charSet,
		defaultLength: randOpts.defaultLength,
		r:             r,
	}, nil
}

// NewID returns a random identifier with the default settings.
func (g *Generator) NewID() string {
	out := make([]rune, g.defaultLength)
	for i := range out {
		out[i] = g.charSet[g.r.Intn(len(g.charSet))]
	}

	return string(out)
}

// NewIDWithLength returns a random identifier with the specified length.
func (g *Generator) NewIDWithLength(l int) (string, error) {
	if l < minLength {
		return "", fmt.Errorf("given ID length is too small, needs to be at least %d characters long", minLength)
	}

	out := make([]rune, l)
	for i := range out {
		out[i] = g.charSet[g.r.Intn(len(g.charSet))]
	}

	return string(out), nil
}
