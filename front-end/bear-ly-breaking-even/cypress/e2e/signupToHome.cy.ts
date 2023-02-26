describe('My First Test', () => {
  it('Visits the signup page', () => {
    cy.visit("http://localhost:3000/signup")

    cy.get("[data-cy=logo]").click()

    // should be redirected to new url without tail /signup
    cy.url().should("not.include", "/signup")

  })
})