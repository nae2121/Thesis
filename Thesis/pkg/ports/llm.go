package ports

type LLM interface {
	Generate(system, user string) (string, error)
}