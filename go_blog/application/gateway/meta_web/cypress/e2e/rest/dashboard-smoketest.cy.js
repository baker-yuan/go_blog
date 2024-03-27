/* eslint-disable no-undef */

context('dashboard page smoke test', () => {
  const selector = {
    pageContent: '.ant-pro-page-container',
  };
  beforeEach(() => {
    cy.login();
  });

  it('should visit dashboard page', function () {
    cy.visit('/');
    cy.contains('Dashboard').click();
    cy.url().should('contains', '/dashboard');
    cy.get(selector.pageContent)
      .children()
      .should('contain', 'Dashboard')
      .and('contain', 'You have not configured Grafana')
      .and('contain', 'Configure');
  });
});
