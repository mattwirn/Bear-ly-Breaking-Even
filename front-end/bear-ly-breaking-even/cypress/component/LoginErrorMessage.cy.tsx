import LoginErrorMessage from "../../components/LoginErrorMessage"

describe('Login error handling', () => {
  it('displays text notifying user that username or password is invalid', () => {
    cy.mount(<LoginErrorMessage exists={false}/>)
    cy.get('div').contains('Invalid username or password.')
  })
  it('returns null if username and password are valid', () => {
    cy.mount(<LoginErrorMessage exists={true}/>)
    cy.get('div').should("be.empty")
  })
})