# gblibrary

> Get book details using ISBN number and build a library page.

### Install
  go get github.com/skippednote/gblibrary

### Setup
- Add a `books.yml` file where you want to build the library page.
```yaml
- isbn: '1617292230'
  rating: 3
  read: false

- isbn: '0137081073'
  rating: 3
  read: false
```
- Add a `books.tmpl.html` which is used to build a page for each book.
```twig
<h1>{{ .Title }}</h1>
<p>{{ .Author }}</p>
```

### Usage
  $ gblibrary

### Usage
- [] Use flags to add custom configuration.
- [] Multiple output formats.
