/* eslint-disable no-undef */

context('Switch language', () => {
  const timeout = 1000;

  const selector = {
    languageSwitcher: '.ant-space-align-center',
  };

  beforeEach(() => {
    cy.login();
  });

  it('should switch language', function () {
    cy.visit('/');

    cy.get(selector.languageSwitcher).click('right');
    cy.contains('简体中文').click({
      force: true,
      timeout,
    });
    cy.contains('服务').click();

    cy.get(selector.languageSwitcher).click('right');
    cy.contains('English').click({
      force: true,
      timeout,
    });
    cy.contains('Create').should('exist');
  });
});
