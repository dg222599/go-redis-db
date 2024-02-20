package command

type Command struct {
	Name  string
	Key   string
	Value interface{}
}

func InitCommand() Command {
	initialCmd := Command{
		Name:  "pseudo",
		Key:   "default",
		Value: 1,
	}

	return initialCmd
}

func ValidateCommand(cmd *Command) (bool, error) {

	return true, nil
}
