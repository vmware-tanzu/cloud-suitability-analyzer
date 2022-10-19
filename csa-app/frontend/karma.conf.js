// Karma configuration file, see link for more information
// https://karma-runner.github.io/1.0/config/configuration-file.html

const jasmineSeedReporter = require('karma-jasmine-seed-reporter');
module.exports = function (config) {
  config.set({
    basePath: '',
    frameworks: ['jasmine', '@angular-devkit/build-angular'],
    plugins: [
      require('karma-jasmine'),
      require('karma-chrome-launcher'),
      require('karma-jasmine-html-reporter'),
      require('karma-coverage'),
      require('@angular-devkit/build-angular/plugins/karma'),
      require('karma-spec-reporter'),
      require("karma-junit-reporter"),
      jasmineSeedReporter
    ],
    client: {
      jasmine: {
        // you can add configuration options for Jasmine here
        // the possible options are listed at https://jasmine.github.io/api/edge/Configuration.html
        // for example, you can disable the random execution with `random: false`
        // or set a specific seed with `seed: 4321`
      },
      clearContext: false, // leave Jasmine Spec Runner output visible in browser
    },
    coverageReporter: {
      dir: require('path').join(__dirname, './coverage/cloud-suitability-analyzer'),
      subdir: '.',
      reporters: [{ type: 'html' }, { type: 'text-summary' }],
      check: {
        global: {
          statements: 70,
          branches: 45,
          functions: 50,
          lines: 70
        }
      },
      watermarks: {
        statements: [ 40, 70 ],
        functions: [ 30, 60 ],
        branches: [ 20, 40 ],
        lines: [ 40, 70 ]
      }
    },
    customLaunchers: {
      ChromeHeadless: {
        base: 'Chrome',
        flags: ['--headless', '--disable-gpu', '--no-sandbox', '--remote-debugging-port=9222']
      }
    },
    specReporter: {
      suppressSkipped: true
    },
    reporters: ['progress', 'kjhtml', 'jasmine-seed', 'spec'],
    port: 9876,
    colors: true,
    logLevel: config.LOG_INFO,
    autoWatch: true,
    singleRun: false,
    restartOnFileChange: true,
    browsers: ['ChromeHeadless'], // ['Chrome'],
  });
};
