# gblibrary-builder

> Get book details using ISBN number and build a library page.

### Install
go get github.com/skippednote/gblibrary-builder

### Usage
- Add a books.yml file where you want to build the library page.
```yaml
- isbn: '1617292230'
  rating: 3
  read: false

- isbn: '0137081073'
  rating: 3
  read: false
```
- Run gblibrary-builder
