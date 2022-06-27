// @ts-check
// Protractor configuration file, see link for more information
// https://github.com/angular/protractor/blob/master/lib/config.ts

const { SpecReporter, StacktraceOption } = require('jasmine-spec-reporter');
const path = require('path');
const downloadsPath = path.resolve(__dirname, '../build/downloads');
/**
 * @type { import("protractor").Config }
 */
exports.config = {
  allScriptsTimeout: 11000,
  specs: ['./src/**/*.e2e-spec.ts'],
  capabilities: {
    chromeOptions: {
      args: ['--no-sandbox', '--window-size=1920,1080', '--allow-insecure-localhost', '--disable-dev-shm-usage'],
      prefs: {
        download: {
          prompt_for_download: false,
          default_directory: downloadsPath
        }
      }
    },
    browserName: 'chrome',
  },
  directConnect: true,
  baseUrl: 'http://localhost:4200/',
  framework: 'jasmine',
  jasmineNodeOpts: {
    showColors: true,
    defaultTimeoutInterval: 30000,
    print: function () {},
  },
  onPrepare() {
    require('ts-node').register({
      project: require('path').join(__dirname, './tsconfig.json'),
    });
    jasmine.getEnv().addReporter(
      new SpecReporter({
        spec: {
          displayStacktrace: StacktraceOption.PRETTY,
        },
      })
    );
  },
};
