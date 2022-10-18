### Conduit Connector Template
This is a template project for building [Conduit](https://conduit.io) connectors in Go. It makes it possible to
start working on a Conduit connector in a matter of seconds.

This template includes the following:
* Skeleton code for the connector's configuration, source and destination
* Example unit tests
* A `Makefile` which commonly used targets
* GitHub workflows to build the code and run the tests
* A dependabot setup which checks your dependencies for available updates and merges minor version upgrades automatically
* Issue and PR templates
* A README template

### How to use
* On this repository's main page, click the "Use this template"
* Enter the information about your repository
* Once your repository has been generated, clone it
* After cloning, run `./setup.sh <module name here>` (for example: 
`./setup.sh github.com/awesome-org/conduit-connector-file`)
* (Optional) Set the code owners (in the `CODEOWNERS` file)

With that, you're all set up and ready to start working on your connector! As a next step, we recommend that you 
check out the [Conduit Connector SDK](https://github.com/ConduitIO/conduit-connector-sdk), which is the Go software 
development kit for implementing a connector for Conduit.


### Repository settings
Following is a list of repository settings we recommend having.

#### General/Pull Requests
1. Allow squash merging only.
2. Always suggest updating pull request branches.
3. Allow auto-merge.
4. Automatically delete head branches.

#### Branch protection rules
Protect the default branch using the following rules:
1. Require a pull request before merging. Require approvals.
2. Require status checks to pass before merging. Require branches to be up to date before merging.
3. Specify status checks that are required.
4. Require conversation resolution before merging.
5. Do not allow bypassing the above settings.

#### Actions/General
1. Allow all actions and reusable workflows.
2. Allow GitHub Actions to create and approve pull requests (if dependabot is used and it's configured to automatically
merge pull requests).

### Specification
The `spec.go` file provides a programmatic representation of the configuration options. This is used by the Conduit
server to validate configuration and dynamically display configuration options to end users.

### Configuration
There are 2 types of "configs":
* general configuration (that applies to both, sources and destinations)
* and source/destination specific configs.

General configs should be added to `config/config.go` whereas any source or destination specific configs should be added
to `source/config.go` and `destination/config.go` respectively
