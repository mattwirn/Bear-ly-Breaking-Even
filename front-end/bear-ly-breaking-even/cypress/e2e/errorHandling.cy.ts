import React from "react"
import ErrorMessage from "@/components/ErrorMessage"

describe('template spec', () => {
  it('renders signup page', () => {
    cy.visit('http://localhost:8080/signup')
  })
  
  it('renders error message if username is taken', () => {
    cy.visit('http://localhost:8080/signup')
    cy.get('input[id=nU').type('text')
    cy.get('input[id=nP]').type('1')
    cy.get('input[id=cP]').type('1')
    cy.get('button[id=signUpButton').click()
    cy.get('[id=errorMessage]')
  })

  it('otherwise displays nothing', () => {
    cy.visit('http://localhost:8080/signup')
    cy.get('[id=errorMessage]').should('not.exist')
  })
})