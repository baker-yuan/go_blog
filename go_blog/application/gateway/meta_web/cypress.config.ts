import { defineConfig } from 'cypress';
import * as globby from 'globby';
import 'dotenv/config';
import defaultSettings from './config/defaultSettings';

const DEFAULT_SETTINGS = defaultSettings.overwrite(process.env)

export default defineConfig({
  viewportWidth: 1920,
  viewportHeight: 1080,
  video: true,
  videoUploadOnPasses: false,
  retries: {
    runMode: 3,
    openMode: 0,
  },
  env: {
    ...process.env,
    DEFAULT_SETTINGS,
    SERVE_URL: DEFAULT_SETTINGS.serveUrlMap[process.env.CYPRESS_SERVE_ENV || 'dev']
  },
  e2e: {
    baseUrl: 'http://localhost:8000',
    setupNodeEvents(on, config) {
      // `on` is used to hook into various events Cypress emits
      // `config` is the resolved Cypress config
      on('task', {
        findFile(mask: any) {
          if (!mask) {
            throw new Error('Missing a file mask to search');
          }

          return globby(mask).then((list) => {
            if (!list.length) {
              throw new Error(`Could not find files matching mask "${mask}"`);
            }

            return list[0];
          });
        },
      });

      require('@cypress/code-coverage/task')(on, config);
      require('cypress-localstorage-commands/plugin')(on, config);
      return config;
    },
  },
});
