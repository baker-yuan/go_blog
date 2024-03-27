
context('Save Paginator Status', () => {
  const timeout = 2000;
  const token = localStorage.getItem('token');
  const { SERVE_URL } = Cypress.env();

  const selector = {
    twentyPerPage: '[title="20 / page"]',
    pageList: '.ant-table-pagination-right',
    pageTwo: '.ant-pagination-item-2',
    pageTwoActived: '.ant-pagination-item-2.ant-pagination-item-active',
    paginationOptions: '.ant-pagination-options',
    deleteButton: '.ant-btn-dangerous',
    notification: '.ant-notification-notice-message',
    notificationCloseIcon: '.ant-notification-close-icon',
  };

  const data = {
    serviceName: 'test_service',
    deleteServiceSuccess: 'Delete Service Successfully',
  };

  before(() => {
    cy.clearLocalStorageSnapshot();
    cy.login();
    cy.saveLocalStorage();
  });

  beforeEach(() => {
    cy.restoreLocalStorage();
  });

  it('should create 11 test services', function () {
    cy.visit('/');
    cy.contains('Service').click();

    for (let i = 0; i <= 10; i += 1) {
      cy.request(
        {
          method: 'POST',
          url: `${SERVE_URL}/apisix/admin/services`,
          headers: {
            Authorization: token,
          },
          body: {
            upstream: {
              nodes: { '39.97.63.215:80': 1 },
              timeout: { connect: 6, read: 6, send: 6 },
              type: 'roundrobin',
              pass_host: 'pass',
            },
            enable_websocket: true,
            name: `${data.serviceName}${i}`,
          },
        },
        {
          retryOnStatusCodeFailure: true,
        },
      ).then((res) => {
        expect(res.body.code).to.equal(0);
      });
    }
    cy.get(selector.pageList).should('be.visible');
  });

  it("should save paginator' status", function () {
    cy.visit('/');
    cy.contains('Service').click();

    // Test page status
    cy.get(selector.pageList).should('be.visible');
    cy.get(selector.pageTwo).click();
    cy.get(selector.pageTwoActived).should('exist');
    cy.location('href').should('include', 'page=2');

    cy.reload();
    cy.get(selector.pageTwoActived).should('exist');
    cy.location('href').should('include', 'page=2');

    // Test pageSize status
    cy.get(selector.paginationOptions).click();
    cy.contains('20 / page').should('be.visible').click();
    cy.get(selector.twentyPerPage).should('exist');
    cy.location('href').should('include', 'pageSize=20');

    cy.reload();
    cy.get(selector.twentyPerPage).should('exist');
    cy.location('href').should('include', 'pageSize=20');
  });

  it('should delete test service', function () {
    cy.visit('/service/list?page=1&pageSize=20');
    cy.reload();
    cy.contains('Service List').should('be.visible');
    cy.get(selector.deleteButton, { timeout })
      .should('exist')
      .each(function ($el) {
        cy.wrap($el).click().click({ timeout });
        cy.contains('button', 'Confirm').click({ force: true });
        cy.get(selector.notification).should('contain', data.deleteServiceSuccess);
        cy.get(selector.notificationCloseIcon).click().should('not.exist');
      });
  });
});
