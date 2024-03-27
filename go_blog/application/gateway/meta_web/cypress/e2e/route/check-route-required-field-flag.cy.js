/* eslint-disable no-undef */

context('Check Route Required Field Flag', () => {
  beforeEach(() => {
    cy.login();
  });

  it('should exist required flag for Route name', function () {
    cy.visit('/');
    cy.contains('Route').click();
    cy.contains('Create').click();
    cy.get('label[title="Name"]').then(($els) => {
      // get Window reference from element
      const win = $els[0].ownerDocument.defaultView;
      // use getComputedStyle to read the pseudo selector
      const before = win.getComputedStyle($els[0], 'before');
      // read the value of the `content` CSS property
      const contentValue = before.getPropertyValue('content');
      // the returned value will have double quotes around it, but this is correct
      expect(contentValue).to.eq('"*"');
    });
  });

  it('should exist required flag for Route path', function () {
    cy.visit('/');
    cy.contains('Route').click();
    cy.contains('Create').click();
    cy.get('label[title="Path"]').then(($els) => {
      const win = $els[0].ownerDocument.defaultView;
      const before = win.getComputedStyle($els[0], 'before');
      const contentValue = before.getPropertyValue('content');
      expect(contentValue).to.eq('"*"');
    });
  });
});
