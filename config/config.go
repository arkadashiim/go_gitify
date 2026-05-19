package config

type Config struct {
	RepoPath string

	GitAuthorName  string
	GitAuthorEmail string

	SSHKeyPath string

	GraphBranch string
}
