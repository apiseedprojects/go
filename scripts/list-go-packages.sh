#!/bin/sh

go list ./... | \
    egrep -v '^github.com/apiseedprojects/go/vendor/' | \
    egrep -v '^github.com/apiseedprojects/go$' | \
    sed 's/github.com\/apiseedprojects\/go\//\.\//g'
