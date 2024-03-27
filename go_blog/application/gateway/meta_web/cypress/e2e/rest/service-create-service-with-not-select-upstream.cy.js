/* eslint-disable no-undef */

context('Edit Service with not select Upstream', () => {
  const selector = {
    name: '#name',
    description: '#desc',
    nodes_0_host: '#submitNodes_0_host',
    nodes_0_port: '#submitNodes_0_port',
    nodes_0_weight: '#submitNodes_0_weight',
    input: ':input',
    notification: '.ant-notification-notice-message',
    nameSearch: '[title=Name]',
    notificationCloseIcon: '.ant-notification-close-icon',
  };

  const data = {
    serviceName: 'test_service',
    createServiceSuccess: 'Create Service Successfully',
    deleteServiceSuccess: 'Delete Service Successfully',
    editServiceSuccess: 'Configure Service Successfully',
    port: '80',
    weight: 1,
    description: 'desc_by_autotest',
    ip1: '127.0.0.1',
    ip2: '127.0.0.2',
    port0: '7000',
    weight0: '1',
  };

  beforeEach(() => {
    cy.login();
  });

  it('should create a test service', function () {
    cy.visit('/');
    cy.contains('Service').click();
    cy.contains('Create').click();
    cy.get(selector.name).type(data.serviceName);
    cy.get(selector.description).type(data.description);
    cy.get(selector.nodes_0_host).click();
    cy.get(selector.nodes_0_host).type(data.ip1);
    cy.get(selector.nodes_0_port).clear().type(data.port0);
    cy.get(selector.nodes_0_weight).clear().type(data.weight0);
    cy.contains('Next').click();
    cy.contains('Next').click();
    cy.get(selector.input).should('be.disabled');
    cy.contains('Submit').click();
    cy.get(selector.notification).should('contain', data.createServiceSuccess);
  });

  it('should edit the service', function () {
    cy.visit('/service/list');

    cy.get(selector.nameSearch).type(data.serviceName);
    cy.contains('Search').click();
    cy.contains(data.serviceName).siblings().contains('Configure').click();
    cy.wait(500);
    cy.get(selector.nodes_0_host).should('not.be.disabled').clear().type(data.ip2);
    cy.get(selector.nodes_0_port).type(data.port);
    cy.get(selector.nodes_0_weight).type(data.weight);
    cy.contains('Next').click();
    cy.contains('Next').click();
    cy.get(selector.input).should('be.disabled');
    cy.contains('Submit').click();
    cy.get(selector.notification).should('contain', data.editServiceSuccess);
  });

  it('should delete this service and upstream', function () {
    cy.visit('/service/list');
    cy.get(selector.nameSearch).type(data.serviceName);
    cy.contains('Search').click();
    cy.contains(data.serviceName).siblings().contains('Delete').click();
    cy.contains('button', 'Confirm').click();
    cy.get(selector.notification).should('contain', data.deleteServiceSuccess);
    cy.get(selector.notificationCloseIcon).click();
  });
});
