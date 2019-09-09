package cmd

// EntityRepository is interface of repositories
type EntityRepository interface {
	Execute() error
}
