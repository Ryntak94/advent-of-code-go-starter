# advent-of-code-go-starter

A starter for our advent of code challenges

For each new challenge, create a folder and a new file under the `main` package.

You can use `template/template.go` as a template.

Make sure to put your input file in the `inputs/` directory.

If you use the `utils.FileAsString()` function, it only needs the file name, and it will open that file from the `inputs/` directory.

To run your file, from the root directory of the repo run `go run [dir]/[file]`, for example, `go run calibration/calibration.go`
It will break if you do it from a different directory as `os.Getwd()` uses the location from which you run the file, not the location of the file.

If you want to be able to test your algorithm, you can use the `test.AssertTestInput` function, passing it your test input string, the function that you are writing for the challenge, and your expected output. You can see `other/other.go` as an example.
