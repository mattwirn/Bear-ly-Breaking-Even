import React from 'react';

describe('template spec', () => {
  it('reveals input fields when edit buttons are clicked', () => {
    cy.visit('http://localhost:3000/dashboard')
    cy.get('button[id=inputButton]').click()
    cy.get('button[id=expensesButton]').click()
  })
  it()
})