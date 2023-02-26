describe('My First Test', () => {
  it('Visits the signup page', () => {
    cy.visit("http://localhost:3000/login")

    cy.get("[data-cy=logo]").click()

    // should be redirected to new url with tail /signup
    cy.url().should('not.include', '/login')

  })
})