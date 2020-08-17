# CSA UI

## Pre-Requisites

- Cloud Suitiability Analyzer

> `CSA UI` was once embedded in `Cloud Suitability Analyzer`. When that tool was open sourced, `CSA UI` was removed and made into a proprietary tool. For the time being, `CSA UI` requires a portion of the OSS project to access the SQLite data stored in the `data/` directory in this repo. A `go` script (`rules-admin.go`)
is used to create the environment `CSA UI` requires to run.

To build and run properly, a fresh clone of `Cloud Suitiabilty is required. 

Follow these instructions:

- Install `go` <https://golang.org/dl/>
- After the install you should have a directory in your home directory: `go/src`

- `cloud-suitability-analyzer` has to be cloned into the directory path `go/src/github.com/vmware-samples/`, so first create this directory path.

- Now, `cd` into that directory and run the following command:

```bash
 git clone https://github.com/vmware-samples/cloud-suitability-analyzer
```

- After the clone:

```bash
cd cloud-suitability-analyzer
git submodule add git@gitlab.eng.vmware.com:app-modernization/csa-internal.git frontend
cd frontend
```


 

- Node
  - Install latest Node LTS version for your OS:

    <https://nodejs.org/en/>
- Watcher

  - required for running tests locally.

```bash
    `brew install watcher`
```

- `concurrently`

```bash
    npm install -g concurrently
```

## To View Available NPM Commands

```bash
npm run
```

## Local Development

1. Run `npm ci` in `../frontend/*` directory
2. Run `npm run dev-start-frontend-backend`
3. Access `localhost:3000` in browser

## Build

```bash
npm run production-build
```

>**Note:** build assets created in `build` directory

## Testing

### To Test Locally

>**Note:** To run tests locally there is a dependency on `watcher`. If you do not have it installed...

```bash
brew install watcher
```

### Run/Watch All Tests

```bash
npm run test -
```

### Run/Watch Single Test

```bash
npm run test 'src/shared/ScoreCard.test.js'
```

### To Test in CI/CD  

```bash
CI=true npm run test
```

### To Generate Test Coverage Report

```bash
npm run test -- --coverage
```

### Audit

#### To Run Audit of NPM Packages

```bash
npm audit
```

#### To Fix Audit of NPM Packages

```bash
npm audit fix
```
