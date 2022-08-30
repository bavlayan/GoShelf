# GoShelf
GoToDo is personal library app that is developed for learning golang.

## Directory Structure
* `bin` - compiled application binaries which will be used for deployment to prod
* `internal` - which includes helper packages used by API and is a specific directory that cannot be imported by code outside of project
* `migrations` - which includes the database migration files
* `remote` - which includes the configuration files to help for production deployment
* `cmd/api` - which includes the application code for our GoShelf API app


## Tips for Golang
* When you designing an API with endpoints in Go, using `http.ServeMux` can be a bit difficult. It doesn't have handlesrs based on HTTP methods (GET, POST, PUT etc.)
* When you are desigining a Rest API, `httprouter` is a solid choice.

### References

Let's Go Further! Advanced patterns for building APIs and web applications in Go | [Alex Edvards](https://lets-go-further.alexedwards.net/)