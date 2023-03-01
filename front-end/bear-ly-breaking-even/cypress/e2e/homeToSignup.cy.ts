describe('My First Test', () => {
    it('Visits the signup page', () => {
      cy.visit("http://localhost:3000")
  
      cy.contains('Sign up').click()
  
      // should redirect router to new url with tail /signup
      cy.url().should('include', '/signup')
  
    })
  })