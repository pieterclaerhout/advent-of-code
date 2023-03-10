package day06

type Command struct{}

func (cmd *Command) Execute(input string) (any, any) {
	return cmd.firstStartOfPackage(input, 4), cmd.firstStartOfPackage(input, 14)
}

func (cmd *Command) firstStartOfPackage(input string, length int) int {

	for i := length; i <= len(input); i++ {
		m := map[rune]struct{}{}
		for _, r := range input[i-length : i] {
			m[r] = struct{}{}
		}

		if len(m) >= length {
			return i
		}
	}

	return -1
}
