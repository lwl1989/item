# item #

go functions about array、map liked php array_values array_columns actions

go version must gte 1.18

[![Dongle Release](https://img.shields.io/github/release/lwl1989/item.svg)](https://github.com/lwl1989/item/releases)
[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/lwl1909/item)
![License](https://img.shields.io/github/license/lwl1909/item)

### Installation

```go
// By github
import (
"github.com/lwl1989/item"
)
```

### Introduction

#### Compare interface

```go
// go comparable
type comparable interface{ comparable }
// item Compare
type ICompare interface {
Compare(ICompare) bool
GetCompareValue() interface{}
}
```

#### map

```go 
MapKeys(mp map[string]T) []string
MapGetKey[T comparable](mp map[string]T, item T) (key string)
MapGetKeyWithCompare[T ICompare](mp map[string]T, item T)  string
```

#### array

```go 
InArray[T comparable](params []T, value T) bool
ArrayUnique[V comparable](params []V) []V 
GetArrayIndex[T comparable](params []T, value T) int 
InArrayCompare[T ICompare](params []T, value T) bool
ArrayMap(params []map[string]interface{}, key string) map[string]map[string]interface{}
ArrayMapCompare[T comparable](params []map[T]interface{}, key T)  map[T]map[T]interface{}
ArrayMapCompareValue[K comparable,V any](params []map[K]V, key K)  map[K]map[K]V
ArrayColumns[K comparable,V any](params []map[K]V, key K) []V
ArrayColumnValues[K comparable,V any](params []map[K]interface{}, key, VKey K) map[K]V
```

#### struct

```go
GetItemMap[K comparable, V any, Item any](values []Item, fieldK, fieldV string) map[K]V
```

### Usage and example