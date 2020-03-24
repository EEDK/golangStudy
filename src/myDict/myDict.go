package myDict

import "errors"

//Dictionary 타입
type Dictionary map[string]string // alias

var (
    errNotFound = errors.New("Not Found")
    errWordExists = errors.New("That word already exists")
    errCanUpdate = errors.New("Cant Update non-exisiting word")
 )

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

 // Update a word
 func (d Dictionary) Update(word , definition string) error {
     _, err := d.Search(word)
     switch err {
     case nil:
         d[word] = definition
     case errNotFound:
         return errCanUpdate
     }
     return nil
 }

//Delete a word
func (d Dictionary) Delete(word string) {
    delete(d, word)
}

 type Money int