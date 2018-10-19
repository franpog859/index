## Contribution rules

If you are a contributor, follow these basic rules:

* The contribution workflow bases on the principles of the [GitHub Flow](https://guides.github.com/introduction/flow/). Thus, the `master` branch is the most important one. Avoid working directly on it. When you work on new features or bug fixes, work on separate branches.
* Work on forks.
* Squash and rebase every pull request before merging.
* You can merge a pull request if you receive an approval from at least one code owner.

Every contributor commits to the following agreement:

* In each pull request, include a description or a reference to a detailed description of the steps that the maintainer goes through to check if a pull request works and does not break any other functionality.
* Provide clear and descriptive commit messages.
* Follow the squash and rebase process.
* Follow the accepted documentation rules and use appropriate templates.
* Choose the proper directory for the content of your pull request.
* As the creator of the pull request, you are responsible for ensuring that the pull request follows the correct review and approval flow.
* Make sure that the person who owns the documentation, such as a technical writer, reviews the changes in the documentation before you merge the pull request.

## Contribution process

This section explains how you can contribute code or content, propose an improvement, or report a bug.

### Contribute code or content

To contribute code or content, follow these steps:

1. Make sure that the change is valid and approved. If you are an external contributor, **open a GitHub issue** before you make a contribution.
2. Fork the repository.
3. Clone it locally, add a remote upstream repository for the original repository, and set up the `master` branch to track the remote `master` branch from the upstream repository.
4. Create a new branch out of the local `master` branch of the forked repository.
5. Commit and push changes to your new branch. Create a clear and descriptive commit message in which you specify what you have changed.
6. Create a pull request from your branch on the forked repository to the `master` branch of the original, upstream repository. Fill in the pull request template according to instructions.
7. If there are merge conflicts on your pull request, squash your commits and rebase the `master` branch.
8. In your pull request:
- Provide a reference to any related GitHub issue.
- Make sure that the [**Allow edits from maintainers**](https://help.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) option is selected to allow upstream repository maintainers, and those with the push access to the upstream repository, to commit to your forked branch.
9. Run the relevant CI job.
10. Wait for the maintainers to review and approve your pull request. The maintainers can approve it, request enhancements to your change, or reject it.

> **NOTE:** The reviewer must check if the related CI job and tests have completed successfully before approving the pull request.

11. When the maintainers approve your change, merge the pull request.

### Report an issue

If you find a bug to report or you want to propose a new feature, create an issue.

> **NOTE:** The repository maintainers handle only well-documented, valid issues that have not been reported yet. Before you create one, check if there are no duplicates. Provide all details and include examples. When you report a bug, list the exact steps necessary to reproduce it.
