import React from 'react';

describe('renders the dashboard page', () => {
  it('passes', () => {
    cy.visit('http://localhost:3000/dashboard')
    cy.get('button')
    cy.get('input[id=h&u]').type('25000');
    cy.get('input[id=transp]').type()
    cy.get('input[id=food]')
    cy.get('input[id=enter]')
    cy.get('input[id=health]')
    cy.get('input[id=other]')
    
  })
})