_This package was developed by [Silicon Ally](https://siliconally.org) while
working on a project for  [Adventure Scientists](https://adventurescientists.org).
Many thanks to Adventure Scientists for supporting [our open source
mission](https://siliconally.org/policies/open-source/)!_

# idgen

[![GoDoc](https://pkg.go.dev/badge/github.com/Silicon-Ally/idgen?status.svg)](https://pkg.go.dev/github.com/Silicon-Ally/idgen?tab=doc)

`idgen` is a simple, zero-dependency Go library for generating random
identifiers. It is useful in places where one might want a UUID-like
identifier, but want to customize length or allowed characters. 

## Usage

```golang

import "github.com/Silicon-Ally/idgen"

...
// Use a cryptographically secure rand source for sensitive applications 
r := rand.New(rand.NewSource(12345))

// Initialize generator with your ID requirements...
generator, err := idgen.New(r, idgen.WithDefaultLength(12), idgen.WithCharSet("abcdef0123456789")) 
if err != nil {
    return nil, fmt.Error("initializing id generator: %w", err)
}

// Generate IDs!
idA := g.NewID()
idB := g.NewID()
...
```
