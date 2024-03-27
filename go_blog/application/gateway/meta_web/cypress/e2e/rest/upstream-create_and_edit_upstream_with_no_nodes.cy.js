/* eslint-disable no-undef */

context('Create and Delete Upstream', () => {
  const selector = {
    name: '#name',
    description: '#desc',
    upstreamNodeMinus0: '[data-cy=upstream-node-minus-0]',
    notification: '.ant-notification-notice-message',
    nameSelector: '[title=Name]',
  };

  const data = {
    upstreamName: 'test_upstream',
    description: 'desc_by_autotest',
    createUpstreamSuccess: 'Create Upstream Successfully',
    configureUpstreamSuccess: 'Configure Upstream Successfully',
    deleteUpstreamSuccess: 'Delete Upstream Successfully',
  };

  beforeEach(() => {
    cy.login();
  });

  it('should create upstream with no nodes', function () {
    cy.visit('/');
    cy.contains('Upstream').click();
    cy.contains('Create').click();

    cy.get(selector.name).type(data.upstreamName);
    cy.get(selector.description).type(data.description);

    // delete all nodes
    cy.get(selector.upstreamNodeMinus0).click();

    cy.contains('Next').click();
    cy.contains('Submit').click();
    cy.get(selector.notification).should('contain', data.createUpstreamSuccess);
    cy.url().should('contains', 'upstream/list');
  });

  it('should configure the upstream with no nodes', function () {
    cy.visit('/');
    cy.contains('Upstream').click();

    cy.get(selector.nameSelector).type(data.upstreamName);
    cy.contains('Search').click();
    cy.contains(data.upstreamName).siblings().contains('Configure').click();
    cy.get(selector.upstreamNodeMinus0).should('not.exist');
    cy.contains('button', 'Next').should('not.be.disabled').trigger('mouseover').click();
    cy.contains('Submit').click({
      force: true,
    });

    cy.get(selector.notification).should('contain', data.configureUpstreamSuccess);
    cy.url().should('contains', 'upstream/list');
  });

  it('should delete the upstream', function () {
    cy.visit('/');
    cy.contains('Upstream').click();
    cy.contains(data.upstreamName).siblings().contains('Delete').click();
    cy.contains('button', 'Confirm').click();
    cy.get(selector.notification).should('contain', data.deleteUpstreamSuccess);
  });
});
