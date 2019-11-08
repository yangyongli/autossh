module autossh

go 1.13

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/pkg/sftp v1.10.1
	golang.org/x/crypto v0.0.0-20191107222254-f4817d981bb6
	golang.org/x/net v0.0.0-20191105084925-a882066a44e0
	golang.org/x/sys v0.0.0-20191105231009-c1f44814a5cd // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
)

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190813064441-fde4db37ae7a
