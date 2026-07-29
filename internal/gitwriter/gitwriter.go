package gitwriter

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/arkadashiim/go_gitify/internal/config"
)

const DEFAULT_COMMIT_MESSAGE = "default"

type GitWriter struct {
	config config.Config
}

func NewGitWriter(config config.Config) *GitWriter {
	return &GitWriter{config: config}
}

func (w *GitWriter) EmptyCommit(date time.Time, message string) ([]byte, error) {
	message = strings.TrimSpace(message)

	if message == "" {
		message = DEFAULT_COMMIT_MESSAGE
	}

	output, err := w.run(
		w.buildCommitEnv(date),
		"commit",
		"--allow-empty",
		"-m",
		message,
	)

	if err != nil {
		return nil, err
	}

	return output, nil
}

func (w *GitWriter) Push() ([]byte, error) {
	output, err := w.run(w.buildPushEnv(), "push", "origin", w.config.GraphBranch)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (w *GitWriter) ForcePush() ([]byte, error) {
	output, err := w.run(w.buildPushEnv(), "push", "--force", "origin", w.config.GraphBranch)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// ResetHistory заменяет ветку GraphBranch на единственный пустой коммит,
// стирая всю предыдущую историю. Локально - push нужно делать отдельно.
func (w *GitWriter) ResetHistory() ([]byte, error) {
	const tmpBranch = "gitify-reset-tmp"

	var out []byte

	if output, err := w.run(nil, "checkout", "--orphan", tmpBranch); err != nil {
		return nil, err
	} else {
		out = append(out, output...)
	}

	if output, err := w.run(nil, "reset", "--hard"); err != nil {
		return nil, err
	} else {
		out = append(out, output...)
	}

	if output, err := w.run(w.buildBaseEnv(), "commit", "--allow-empty", "-m", "reset"); err != nil {
		return nil, err
	} else {
		out = append(out, output...)
	}

	if output, err := w.run(nil, "branch", "-M", tmpBranch, w.config.GraphBranch); err != nil {
		return nil, err
	} else {
		out = append(out, output...)
	}

	return out, nil
}

func (w *GitWriter) run(
	env []string,
	args ...string,
) ([]byte, error) {
	cmd := exec.Command("git", args...)

	cmd.Dir = w.config.RepoPath
	cmd.Env = append(os.Environ(), env...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf(
			"git %v failed: %w\n%s",
			args,
			err,
			output,
		)
	}

	return output, nil
}

func (w *GitWriter) buildBaseEnv() []string {
	return []string{
		"GIT_AUTHOR_NAME=" + w.config.GitAuthorName,
		"GIT_AUTHOR_EMAIL=" + w.config.GitAuthorEmail,

		"GIT_COMMITTER_NAME=" + w.config.GitAuthorName,
		"GIT_COMMITTER_EMAIL=" + w.config.GitAuthorEmail,
	}
}

func (w *GitWriter) buildCommitEnv(date time.Time) []string {
	commitDate := date.Format(time.RFC3339)

	return append(
		w.buildBaseEnv(),

		"GIT_AUTHOR_DATE="+commitDate,
		"GIT_COMMITTER_DATE="+commitDate,
	)
}

func (w *GitWriter) buildPushEnv() []string {
	return append(
		w.buildBaseEnv(),
		"GIT_SSH_COMMAND="+w.sshCommand(),
	)
}

func (w *GitWriter) sshCommand() string {
	return fmt.Sprintf(
		"ssh -i %s -o IdentitiesOnly=yes",
		w.config.SSHKeyPath,
	)
}
