describe("AppComponent", () => {
  it("Search Parameter Test", () => {
    const searchTerm = "cen";

    cy.visit("/");
    cy.get("[data-cy=search]").type(searchTerm).type('{enter}')

    cy.url().should("include", "/results");
    cy.location("pathname").should("eq", `/results/${searchTerm}`);
  });
});
