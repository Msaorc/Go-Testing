# Go-Testing
 Repository to practice and learn testing in Golang.

### Package testing
 * **Errorf**: To show an error in tests

### Commands
 * **To Exec Test**: go test -v
 * **To see the coverage**: go test -coverprofile=coverage.out
 * **To see coverage visually**: go tool cover -html=coverage.out
 * **To Exec Benchmark with tests**: go test -bench=.
 * **For Exec Benchmark with specific tests**: go test -bench=. -run=^#
 * **For Exec Benchmark with timeout**: go test -bench=. -run=^# -benchtime=-3s
 * **To Exec Fuzzing with tests**: go test -fuzz=.
 * **For Exec Fuzzing with specific tests**: go test -fuzz=. -run=^#
 * **For Exec Fuzzing with timeout**: go test -bench=. -run=^# -fuzztime=5s
 * **For more information**: go help test

### Package testify
* **Equal**: To test equality of values
* **Error**: To test return errors
* **Contains**: Check if it contains a value within a value string
* **For more information**: `https://pkg.go.dev/github.com/stretchr/testify@v1.8.1/assert` / `https://github.com/stretchr/testify`
