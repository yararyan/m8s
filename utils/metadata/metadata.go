package metadata

import (
	"strings"
)

const (
	// AnnotationBitbucketBranch is an identifier for Bitbucket branch this environment was built from.
	AnnotationBitbucketBranch = "bitbucket.branch"

	// AnnotationBitbucketRepoOwner is an identifier for Bitbucket repository owner.
	AnnotationBitbucketRepoOwner = "bitbucket.repo.owner"

	// AnnotationBitbucketRepoName is an identifier for Bitbucket repository this environment was built from.
	AnnotationBitbucketRepoName = "bitbucket.repo.name"

	// AnnotationCircleCIRepositoryURL is an identifier for CircleCI repository the environment was built from.
	AnnotationCircleCIRepositoryURL = "circleci.repository.url"

	// AnnotationCircleCIPRNumber is an identifier for CircleCI pull request the environment was built from.
	AnnotationCircleCIPRNumber = "circleci.pr.number"

	// AnnotationCircleCIPRUsername is an identifier for CircleCI pull request submitted by a user.
	AnnotationCircleCIPRUsername = "circleci.pr.username"
)

// Annotations are used for attaching metadata to a environment.
func Annotations(envs []string) (map[string]string, error) {
	annotations := make(map[string]string)

	for _, env := range envs {
		sl := strings.Split(env, "=")

		if len(sl) != 2 {
			continue
		}

		switch sl[0] {
		// Check if we have Bitbucket Pipelines environment variables.
		// https://confluence.atlassian.com/bitbucket/environment-variables-794502608.html
		case "BITBUCKET_BRANCH":
			annotations[AnnotationBitbucketBranch] = sl[1]
		case "BITBUCKET_REPO_OWNER":
			annotations[AnnotationBitbucketRepoOwner] = sl[1]
		case "BITBUCKET_REPO_SLUG":
			annotations[AnnotationBitbucketRepoName] = sl[1]
		// Check if we have CircleCI environment variables.
		// https://circleci.com/docs/2.0/env-vars/
		case "CIRCLE_REPOSITORY_URL":
			annotations[AnnotationCircleCIRepositoryURL] = sl[1]
		case "CIRCLE_PR_NUMBER":
			annotations[AnnotationCircleCIPRNumber] = sl[1]
		case "CIRCLE_PR_USERNAME":
			annotations[AnnotationCircleCIPRUsername] = sl[1]
		}
	}

	return annotations, nil
}
