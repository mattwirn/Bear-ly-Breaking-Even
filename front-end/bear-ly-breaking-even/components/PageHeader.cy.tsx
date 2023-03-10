import React from 'react'
import * as NextRouter from 'next/router'
import PageHeader from './PageHeader'
import Logo from './Logo'

describe('<PageHeader />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<PageHeader />)
  })
  it("dispays the header component", () => {
    cy.mount(<PageHeader />)
    const pathname = '../pages/index.tsx'
    const push = cy.stub()
    cy.stub(NextRouter, 'useRouter').returns({ pathname, push })
  })
})