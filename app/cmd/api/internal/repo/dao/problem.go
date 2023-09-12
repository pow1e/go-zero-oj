package dao

type IProblem interface {
	GetProblemList(keyword string) ([]IProblem, error)
}
