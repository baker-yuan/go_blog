/* eslint-disable no-undef */

context('settings page smoke test', () => {
  const selector = {
    avatar: '.ant-space-align-center',
    pageContainer: '.ant-pro-page-container',
    grafanaURL: '#grafanaURL',
    notificationMessage: '.ant-notification-notice-message',
    explain: '.ant-form-item-explain',
  };

  const data = {
    invalidURL: 'httx://www.test.com',
    validURL: 'http://localhost:8000/routes/list',
    fetchURL: 'fetchURL',
    fetch: '@fetchURL',
    grafanaAddress: 'Grafana Address',
    grafanaExplanation1: 'Grafana address should begin with HTTP or HTTPS',
    grafanaExplanation2: 'Address is invalid',
    updateSuccessfully: 'Update Configuration Successfully',
  };

  beforeEach(() => {
    cy.login();
  });

  it('should visit settings page', function () {
    cy.visit('/');
    cy.get(selector.avatar).invoke('show').click('center');
    cy.contains('Settings').click();
    cy.url().should('contains', '/settings');
    cy.get(selector.pageContainer)
      .children()
      .should('contain', 'Setting')
      .and('contain', data.grafanaAddress)
      .and('contain', data.grafanaExplanation1);
  });

  it('should set a invalid url', function () {
    cy.visit('/');
    cy.get(selector.avatar).invoke('show').click('center');
    cy.contains('Settings').click();
    cy.url().should('contains', '/settings');
    cy.get(selector.grafanaURL).clear().type(data.invalidURL);
    cy.get(selector.explain).should('contain', data.grafanaExplanation2);
  });

  it('should set a accessible URL', function () {
    cy.visit('/');
    cy.get(selector.avatar).invoke('show').click('center');
    cy.contains('Settings').click();
    cy.url().should('contains', '/settings');
    cy.get(selector.grafanaURL).clear().type(data.validURL);
    cy.contains('Submit').click();

    cy.get(selector.notificationMessage).should('contain', data.updateSuccessfully);
    cy.intercept(data.validURL).as(data.fetchURL);
    cy.wait(data.fetch);
    cy.get(selector.pageContainer).children().should('contain', 'Dashboard');
  });
});
