package test

func AssertTestInput[T string | int](testInput string, algorithm func(string) T, expectedOutput T) bool {
	receivedOutput := algorithm(testInput)

	return receivedOutput == expectedOutput
}
