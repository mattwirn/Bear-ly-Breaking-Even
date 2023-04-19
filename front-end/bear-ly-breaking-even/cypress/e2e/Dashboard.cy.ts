describe('template spec', () => {
  it('renders the dashboard', () => {
    
    cy.visit('http://localhost:8080/dashboard/')
  })
  it('presses the expenses button, and displays tables', () => {
    cy.visit('http://localhost:8080/dashboard/')
    cy.get('button[id=ExpButton]').click()

  })
})