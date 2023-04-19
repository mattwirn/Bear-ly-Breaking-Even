import React from 'react'
import Table from '../../components/ExpenseTable'

describe('<Table />', () => {
  it('renders html of table component', () => {
    cy.mount(<Table />)
  })
  it('lets users add rows to the table', () => {
    cy.mount(<Table />)
    cy.get('button[id=addRowButton]').click()
  })
  it('allows users to delete rows from the table', () => {
    cy.mount(<Table />)
    cy.get('button[id=deleteRowButton]').click()
  })
})