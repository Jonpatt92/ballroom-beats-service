# ballroom-beats-service
Golang API to expose Spotify's song analysis API for the Ballroom Beats Mobile App


## Requirements

`dep` -- run `brew install dep`

From within your project repo folder, run `dep init`; this takes several moments to run. If the file `Gopkg.toml` already exists, this will exit with a warning and not do anything new, which is fine.

Then run `go get .`

## Deployment Notes

Heroku needed several files which are added to the root folder of the repo, in order to deploy properly to their platform. This process will create a new binary called `bin/ballroom-blitz` instead of `main` binary which can be deleted from your repo. To compile the binary with this name and path, from the root folder of this project run the following command:

`go build -o bin/ballroom-blitz -v .`

The `bin` folder will be created for you.

Then run the following:

`heroku buildpacks:clear`


You should see output that looks like:

```bash
$ go build -o bin/ballroom-blitz -v .
ballroom-beats-service/vendor/github.com/go-playground/locales/currency
ballroom-beats-service/vendor/github.com/lib/pq/oid
ballroom-beats-service/vendor/github.com/jinzhu/inflection
ballroom-beats-service/vendor/github.com/go-playground/locales
ballroom-beats-service/vendor/github.com/mattn/go-isatty
ballroom-beats-service/vendor/github.com/gin-gonic/gin/internal/json
ballroom-beats-service/vendor/github.com/gin-contrib/sse
ballroom-beats-service/vendor/github.com/golang/protobuf/proto
ballroom-beats-service/vendor/github.com/go-playground/universal-translator
ballroom-beats-service/vendor/github.com/ugorji/go/codec
ballroom-beats-service/vendor/github.com/leodido/go-urn
ballroom-beats-service/vendor/gopkg.in/yaml.v2
ballroom-beats-service/vendor/gopkg.in/go-playground/validator.v9
ballroom-beats-service/vendor/github.com/jinzhu/gorm
ballroom-beats-service/vendor/github.com/lib/pq/scram
ballroom-beats-service/vendor/github.com/lib/pq
ballroom-beats-service/vendor/github.com/gin-gonic/gin/binding
ballroom-beats-service/vendor/github.com/gin-gonic/gin/render
ballroom-beats-service/vendor/github.com/gin-gonic/gin
ballroom-beats-service
```



