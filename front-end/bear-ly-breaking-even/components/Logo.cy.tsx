import React from 'react'
import Logo from './Logo'

describe('<Logo />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<Logo w={40} />)
  })
  it('should show the icon', () => {
    cy.mount(<Logo w={40} />)
    cy.get('[data-cy=logo]').should('be.visible')
  })
})