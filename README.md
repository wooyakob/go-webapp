Go
Writing Web Applications: https://go.dev/doc/articles/wiki/

Create data structure with load and save methods.
Use net/http pkg to build web apps.
Use html/template pkg to process HTML templates.
Use regexp pkg to validate user input.
Use closures.

Check Go version and installation:
jake.wood@D6TGGN7D2F go-webapp % go version
go version go1.25.4 darwin/arm64


How to write go code:
- Development of a simple Go package inside a module
- Go tool to fetch, build, install Go modules, packages and commands
- Go programs are organized into packages
- A package is a collection of source files in the same directory that are compiled together
- Functions, types, variables and constants defined in one source file are visible to all other source files in the same package
- A repository contains one or more modules
- A module is a collection of related Go packages that are released together
- A Go repo typically contains only one module, located at the root of the repo. go.mod
- go.mod declares the module path: import prefix for all packages in the module
- You don't need to publish code to a remote repo before you can build it
- It is a good habit to organize your code as if you will publish it someday
- Each module's path also indicates where the go command should look to download it
- Import path is a string used to import a package
- go mod init example/user/hello
- cat go.mod module example/user/hello
- go install example/user/hello
- GOBIN binaries installed to this directory
- GOPATH binaries are installed to the bin subdirectory of the first directory in the GOPATH list
- example "github.com/google/go-cmp/cmp" dependency on an external module, need to download it and record its version in go.mod file
- go mod tidy adds missing module requirements for imported packages and removes requirments on modules that are not used anymore
- to remove all downloaded modules, go clean -modcache
- testing: lightweight test framework composed of go test command and testing package
- write a test file with a name ending in _test.go that contains functions named TestXXX with signature func (t *testing.T)
- the test runs each such function, if function calls a failure such as t.Error or t.Fail, the test is considered to have failed
- run go test


New tasks:
TODO

Store templates in tmpl/ and page data in data/.

Add a handler to make the web root redirect to /view/FrontPage.

Spruce up the page templates by making them valid HTML and adding some CSS rules.

Implement inter-page linking by converting instances of [PageName] to
<a href="/view/PageName">PageName</a>.

 (hint: you could use regexp.ReplaceAllFunc to do this)



 go build wiki.go 
 ./wiki          
