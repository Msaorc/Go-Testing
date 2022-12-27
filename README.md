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
 * **For more information**: go help test