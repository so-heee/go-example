module github.com/so-heee/golang-example/testing

go 1.13

replace (
	github.com/so-heee/golang-example/testing/pkg/benchmark => ./pkg/benchmark
	github.com/so-heee/golang-example/testing/pkg/calc => ./pkg/calc
	github.com/so-heee/golang-example/testing/pkg/examples => ./pkg/examples
	github.com/so-heee/golang-example/testing/pkg/testmain => ./pkg/testmain
	github.com/so-heee/golang-example/testing/pkg/unorderedoutput => ./pkg/unorderedoutput
)

require (
	github.com/so-heee/golang-example/testing/pkg/benchmark v0.0.0-00010101000000-000000000000
	github.com/so-heee/golang-example/testing/pkg/calc v0.0.0-00010101000000-000000000000
	github.com/so-heee/golang-example/testing/pkg/examples v0.0.0-00010101000000-000000000000
	github.com/so-heee/golang-example/testing/pkg/testmain v0.0.0-00010101000000-000000000000
	github.com/so-heee/golang-example/testing/pkg/unorderedoutput v0.0.0-00010101000000-000000000000
)
