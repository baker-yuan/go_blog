/* eslint-disable no-undef */

context('ssl smoke test', () => {
  const selector = {
    notificationMessage: '.ant-notification-notice-message',
    notificationDesc: '.ant-notification-notice-description',
  };

  const data = {
    deleteSSLSuccess: 'Remove target SSL successfully',
    sslErrorAlert: "key and cert don't match",
  };

  beforeEach(() => {
    cy.login();
    cy.fixture('certificate.json').as('certificate');
  });

  it('should set match certificate and key by input', function () {
    // use `function () if used `fixture` above`
    // go to ssl create page
    cy.visit('/');
    cy.contains('SSL').click();
    cy.contains('Create').click();

    const validCert = this.certificate.valid.cert;
    const validKey = this.certificate.valid.key;
    cy.get('#cert').type(validCert);
    cy.get('#key').type(validKey);

    cy.contains('Next').click();
    cy.contains('Submit').click();
    cy.url().should('contains', 'ssl/list');
  });

  it('should delete the ssl record just created', function () {
    cy.visit('/');
    cy.contains('SSL').click();
    const sni = this.certificate.valid.sni;
    cy.contains(sni).parents().contains('Delete').click();
    cy.contains('button', 'Confirm').click();
    cy.get(selector.notificationMessage).should('contain', data.deleteSSLSuccess);
  });

  it('should set unmatch certificate and key by input', function () {
    // go to ssl create page
    cy.visit('/');
    cy.contains('SSL').click();
    cy.contains('Create').click();

    const invalidCert = this.certificate.invalid.cert;
    const invalidKey = this.certificate.invalid.key;
    cy.get('#cert').type(invalidCert);
    cy.get('#key').type(invalidKey);

    cy.contains('Next').click();
    cy.get(selector.notificationDesc).should('contain', data.sslErrorAlert);
  });
});
