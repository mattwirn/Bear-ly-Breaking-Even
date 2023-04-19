describe('template spec', () => {
  it('visits the login page', () => {
    cy.visit('http://localhost:8080/login')
  })
  it("renders error message if username and password arent in database", () => {
    cy.visit('http://localhost:8080/login')
    cy.get('input[id=username]').type('matt')
    cy.get('input[id=password]').type('2')
    cy.get('button[id=loginButton]').click()
    cy.get('[id=errorMessage]')
  })
  it("redirects user to dashboard if username and password exist", () => {
    cy.visit('http://localhost:8080/login')
    cy.get('input[id=username]').type('matt')
    cy.get('input[id=password]').type('1')
    cy.get('button[id=loginButton]').click()
    cy.url().should('include', '/dashboard')
  })
})