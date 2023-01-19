# Suggestions for improvement

* Provide prebuild binaries for users, do not force them to build the binary themselves
* + Makefile is missing .PHONY declarations, which will cause issues if files exist with the name of the target
* + The first make target is expected to be "all" and to build all binaries of the project
* + The commands which make executes should not be hidden with @
* + The process for building the executables should only be 1) clone the repository 2) run "make". There should be no need
  to manually run generate-client-code, all, test, lint. that should be automated through the make target dependencies.
* The tableprinter produces a lot of clutter in the console which the user has to visually filter through. Established
  best practices for command line tools like docker and kubectl use the go text/tabwriter (part of the go environment) and do
  not try to paint tables as text like excel would show a table.
* + Do not keep around disabled code which is commented out (i.e. .golangci.yaml)
* + Do not rename the executable to /app/main in the Dockerfile
* + Get rid of the githooks - linting can be done with make targets automatically, and you do not want to prevent work in
  progress to be pushed to some feature branch
* There should be a make target to automatically pull the argus openapi spec from some place when updates are needed
* + The default location the cli should look for the configuration file should be the home directory (~/.stackit-argus-cli.yaml)
* + Provide an internal default for base_url, so not every user needs to provide that, as that url is the same for everybody
* + Provide a "configure" command to fill in the config file (aws does this)
* Be careful about what you provide outside the internal directory. You are comitting a public API there which
  you cannot easily change later on. the cmd subdirectory should really only contain the barebones of the cmd, all support
  data types and functions should be in the internal subdirectory. the pkg/argus package I would also move to internal
* + Keep a consistent naming of command line parameters. Stick to "this-is-a-multi-word-parameter" instead of "InstanceId"
  and "projectId", as that naming is already standard from tools like docker and kubectl
* + Do not output status messages when everything is fine. That status output does not provide any additional value
  and will even lead to broken yaml and json output when piped into a file.
* + Do not invent a custom logger interface (internal/log/logger.go). Use established best practices and use logr.Logger
  with a zap backend.
* + Keep to the golang established best-practices for naming subdirectories for packages (use "output_tables" instead of
  "outputTables" and "yaml_to_json" instead of "yamltojson"). the binary can still be named "stackit-argus-cli" as the
  package in that directory is named "main".
