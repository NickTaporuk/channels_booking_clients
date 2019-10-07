package config

// CompilationConfiguration use for set data from compilation variables :
//     version - version of the compiled app
//     commit  - git commit of a a compilation
//     date    - date of a compilation the application
type CompilationConfiguration struct {
	version string
	commit  string
	date    string
}

// Version - getter of version the app
func (c *CompilationConfiguration) Version() string {
	return c.version
}

// SetVersion - setter of version the app
func (c *CompilationConfiguration) SetVersion(version string) {
	c.version = version
}

// Commit - getter of a compilation of commit
func (c *CompilationConfiguration) Commit() string {
	return c.commit
}

// SetCommit - setter of a compilation of commit
func (c *CompilationConfiguration) SetCommit(commit string) {
	c.commit = commit
}

// Date - getter of a compilation date
func (c *CompilationConfiguration) Date() string {
	return c.date
}

// SetDate - setter of a compilation date
func (c *CompilationConfiguration) SetDate(date string) {
	c.date = date
}
