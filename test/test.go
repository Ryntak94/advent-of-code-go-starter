package test

func AssertTestInput[T string | int](testInput string, algorithm func(string) T, expectedOutput any) bool {
	receivedOutput := algorithm(testInput)

	return receivedOutput == expectedOutput
}
