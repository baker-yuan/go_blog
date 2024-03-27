/* eslint-disable no-undef */
import './commands';
import '@cypress/code-coverage/support';
import 'cypress-localstorage-commands';

const { SERVE_ENV = 'dev' } = Cypress.env();

before(() => {
  // reset etcd before test
  if (SERVE_ENV === 'test') {
    cy.exec('etcdctl del --prefix /', { failOnNonZeroExit: false });
  }
});

Cypress.on('uncaught:exception', () => {
  // returning false here prevents Cypress from
  // failing the test
  return false;
});
