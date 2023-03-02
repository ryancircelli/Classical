import { AppComponent } from "./app.component"

describe('AppComponent', () => {
  it('mounts', () => {
    cy.mount(AppComponent)
  })
})