import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { ClassAPIService } from 'src/app/services/class.services';

import { AddClassComponent } from './add-class.component';

function createClassAPIServiceStub(): ClassAPIService {
  return {
    apiUrl: '',
    http: null,
    getClasses: cy.stub(),
    addClass: cy.stub(),
  } as unknown as ClassAPIService;
}

function createRouterStub(): Router {
  return {
    navigate: cy.stub(),
  } as unknown as Router;
}


describe('submitNewClass()', () => {

  let classComponent: AddClassComponent;

  beforeEach(() => {
    classComponent = new AddClassComponent(createClassAPIServiceStub(), createRouterStub());
  });

  it('should set errorMessage when class name is too short or too long', () => {
    classComponent.newClass = 'abc';
    classComponent.submitNewClass().catch(errorMessage => {
      expect(errorMessage).to.equal('Class name must be between 4 and 10 characters long');
    });

    classComponent.newClass = 'abcdefghijk';
    classComponent.submitNewClass().catch(errorMessage => {
      expect(errorMessage).to.equal('Class name must be between 4 and 10 characters long');
    });
  });

  it('should call addClass API and navigate when class name is valid', () => {
    classComponent.newClass = 'test123';
    classComponent.classAPIService.addClass = cy.stub().callsFake(() => Promise.resolve());
    classComponent.router.navigate = cy.stub().callsFake(() => Promise.resolve());

    classComponent.submitNewClass().then(() => {
      expect(classComponent.classAPIService.addClass).to.have.been.calledWith('test123');
      expect(classComponent.router.navigate).to.have.been.calledWith(['/class', 'test123']);
    });
  });

  it('should handle API error when class already exists', () => {
    classComponent.newClass = 'cis4930';
    classComponent.classAPIService.addClass = cy.stub().callsFake(() => Promise.reject({ text: 'Class with Name = cis4930 already exists' }));

    classComponent.submitNewClass().catch(errorMessage => {
      expect(errorMessage).to.equal('Class with Name = cis4930 already exists');
    });
  });

  it('should handle general API errors', () => {
    classComponent.newClass = 'test123';
    classComponent.classAPIService.addClass = cy.stub().callsFake(() => Promise.reject({ message: 'API error' }));

    classComponent.submitNewClass().catch(errorMessage => {
      expect(errorMessage).to.equal('API error');
    });
  });
});
