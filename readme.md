Very low effort UNIX command for extracting the article body from a HTML file.

Install: clone this repo then `go install .`  

Usage:

```
cat somefile.html | iconv -f gbk | readability-cli
# iconv is for converting encoding to utf-8; it's optional
```
