import React from 'react';
import ErrorMessage from '../../components/ErrorMessage';

describe('ErrorMessage.cy.tsx', () => {
  it('notifies user if username is already taken', () => {
    cy.mount(<ErrorMessage taken={true}/>)
    cy.get('div').contains('The username you entered is taken. Please choose another.')
  })
  it('returns null if username is not taken', () => {
    cy.mount(<ErrorMessage taken={false}/>)
    cy.get('div').should("be.empty")
  })

})