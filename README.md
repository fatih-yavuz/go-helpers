# go-helpers
Utility library for personal use. I'll keep adding helper functions as I need. Feel free to contribute.

## bytes
    func Compress(byt []byte) ([]byte, error)

## function
    func Caller() string
    func Name(i interface{}) string

## log
    func Error(err error)

## strings
    func RemoveAfter(str string, after string) string
    func RemoveAfterwards(str string, after string) string 
### extract
#### url
##### from
    func FilePath(path string) (string, error)
### normalize
    func Url(path string) (string, error)
   

## url
    func FilePath(u string) (string, error)
    func Host(u string) (string, error)
    func Validate(u string) bool
