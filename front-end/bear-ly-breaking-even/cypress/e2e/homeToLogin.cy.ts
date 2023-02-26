describe('My First Test', () => {
  it('Visits the signup page', () => {
    cy.visit("http://localhost:3000")

    cy.contains('Log in').click()

    // should redirect router to new url with tail /login
    cy.url().should('include', '/login')

  })
})