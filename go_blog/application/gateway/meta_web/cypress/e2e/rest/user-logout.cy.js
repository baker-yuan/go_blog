/* eslint-disable no-undef */

context('Logout Test', () => {
  beforeEach(() => {
    cy.login();
  });

  it('logout', () => {
    cy.visit('/');
    cy.contains('.anticon', 'APISIX User', {
      matchCase: false,
    }).click({
      force: true,
    });
    cy.get('[aria-label=logout]').click();
    cy.url().should('contains', '/user/login');
  });
});
