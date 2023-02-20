import React from 'react'
import Logo from './Logo'

describe('<Logo />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<Logo w={40} />)
  })
  it('should render the icon', () => {
    cy.mount(<Logo w={100} />)
    cy.get('img').should('be.visible')
  })
})