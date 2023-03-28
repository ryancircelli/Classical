import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

import { ClassAPIService } from './class.services'

describe('ClassAPIService', () => {

  let classAPIService: ClassAPIService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [ClassAPIService]
    });

    classAPIService = TestBed.inject(ClassAPIService);
    httpMock = TestBed.inject(HttpTestingController);
   
    // let apiUrl: string = 'http://localhost:8000';

    // cy.intercept('GET', `${apiUrl}/getClasses`, (req) => {
    //   req.reply(classes);
    // }).as('backendAPI');
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('test', () => {
    let apiUrl: string = 'http://localhost:8000';

    cy.intercept('GET', `${apiUrl}/getClasses`).as('getClasses');
    const mockResponse = [{ id: 1, className: 'Class 1' }, { id: 2, className: 'Class 2' }];
    const request = classAPIService.getClasses();
    cy.wait('@getClasses');
    cy.wrap(request).then((observable) => {
      const req = httpMock.expectOne(`${apiUrl}/getClasses`);
      req.flush(mockResponse);
      observable.subscribe((response) => {
        expect(response).to.deep.equal(mockResponse);
      });
    });
  });
});