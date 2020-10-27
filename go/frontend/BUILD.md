# NOTE: This repo has been copied to it's new home in preparation for becoming open source.

# Cloud Suitability Analyzer (CSA)

This tool allows data collection and visualization functionality to be housed in a native-code executable. More importantly, it will allow `forge` to function as a remote probe, collecting legacy meta-data and allowing customers to easily return that data for further analysis.

## NOTE:

_`CSA` has been adopted by the `Journey` team and its development will become a formal activity within the organization with Steve Woods (swoods@vmware.com) leading the effort full-time. Our first priority is to more completely document the source code and to produce a contributors guide for rules and for source code, `Go` and `React`. The local development instructions (below) have been corrected_

## License

`CSA`, and `forge`are to be licensed under an internal license which is under
consideration now.

## Developing

To get started developing, you'll need to clone this project inside your `GOPATH`. If you haven't explicitly set this, it will default to `$HOME/go`.

Cloning the project should be inside a path matching the GitLab path:

If the parent directories don't exist, you'll need to make them.

## Important: If you intend to commit changes, you MUST use the exact command below. Our build is very sensitive to `GOPATH`.

```
git clone https://gitlab.eng.vmware.com/app-modernization/cloud-suitability-analyzer
```

This uses `git clone` instead of `go get` because `go get` does not allow choosing SSH or HTTPS - it always assumes HTTPS.
It's possible to force `go get` to use SSH, by adding the following to your `.gitconfig`:

```
[url "git@gitlab.eng.vmware.co:"]
  insteadOf = https://gitlab.eng.vmware.com
```

Once checked out, run `make deps` to fetch dependencies, then `make gen` to generate required code.

To test, you need to do one of the following:

### Separate backend/frontend

This will run the backend and the frontend as separate processes, making the development cycle a bit quicker.

See the instructions in [`frontend/README.md#Local Development`](frontend/README#Local Development) for more details.

### Single binary

This will produce a production-ready binary with all frontend resources baked in.

We use the `go-bindata` tool to embedded web application resources into our `Go` executable.

This package converts any file into managable `Go` source code. Useful for embedding binary data into a go program.

You'll need to install [go-bindata](https://github.com/go-bindata/go-bindata)
with the command:

```bash
 go get -u github.com/go-bindata/go-bindata/...
```

Then you can perform the bind with the following command run from the `go/frontend` directory:

```bash
${GOPATH}/bin/go-bindata -o ../resources/admin-site.go -pkg resources build/...
```

Then you can build the binary using the following commands run from the `go` directory:

- Update dependencies

```bash
make deps
```

- Generate code such as the code required by rules

```bash
make gen
```

- Build the `forge` executable. This command will differ on `Linux` and `Windows`

```bash
GOOS=darwin GOARCH=amd64 go build -o ./exe/forge forge.go
```

## Documentation tools

`PAA` is a complex, highly concurrent application that fully utilizes the `Go` concurrency paradigm. As such it is helpful to leverage `Go` toolsets

### `go-callvis`

`go-callvis` produces a useful visualization of `forge` call graphs. It is helpful to understand the anatomy of an application. Follow these steps to use the tools:

1. Install [go-callvis](https://github.com/TrueFurby/go-callvis)
2. Run this command:

```
go-callvis -nostd -group pkg,type
~/go/src/gitlab.eng.vmware.co/go
```

![CallGraph](misc/images/app-analyzer-CallGraph.png)

If you are going to be working in the `Go` code, we have some scripts that will de-clutter
the diagram. These can be found in the `misc/bin` directory, which you might want to
add to your path. The scripts details:

- `visAll.sh`: All `Go` code
- `visBackend.sh`: Backend code
- `visDb.sh`: Database code
- `visAppAnalyzer:` Core parsing and rule processing code

Note: The scripts will take upto 30 seconds to finish. The end results should launch
a browser that allows you to drill down and explore the code. Just `<Ctrl> C` to exit
the script.

## Usage

For detailed usage instructions see the manual that is contained in the download.
