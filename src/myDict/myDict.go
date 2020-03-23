package myDict

import "errors"

//Dictionary 타입
 type Dictionary map[string]string // alias
var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")

 //Search for a word
 func (d Dictionary) Search(word string)(string, error) {
  value , exists := d[word]
  if exists {
   return value, nil
  }
  return "" , errNotFound
 }

 func (d Dictionary) Add(word, def string) error {
   _ , err := d.Search(word)
   if err == errNotFound {
    d[word] = def
   } else if err == nil {
    return errWordExists
   }
   return nil
 }

 type Money int

