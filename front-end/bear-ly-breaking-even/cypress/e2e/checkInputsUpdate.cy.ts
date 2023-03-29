import React from 'react';

describe('template spec', () => {
  
  it('reveals input fields when edit buttons are clicked', () => {
    cy.visit('http://localhost:3000/dashboard')
    cy.get('button[id=inputButton]').click()
    cy.get('button[id=expensesButton]').click()
  })

  it('inputs new values for each input field, and submit new values', () => {
    cy.visit('http://localhost:3000/dashboard')
    cy.get('button[id=inputButton]').click()
    cy.get('button[id=expensesButton]').click()
    cy.get('input[id="h&u"]').type('100')
    cy.get('input[id=transp]').type('100')
    cy.get('input[id=food]').type('100')
    cy.get('input[id=enter]').type('100')
    cy.get('input[id=health]').type('100')
    cy.get('input[id=other]').type('100')
    cy.get('button[id=expensesButton]').click()

  })
})