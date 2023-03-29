import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { ClassAPIService } from './class.services';
import { Class } from '../types';

describe('ClassAPIService', () => {
  let classAPIService: ClassAPIService;
  let httpTestingController: HttpTestingController;
  const apiUrl = 'http://localhost:8000';

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [ClassAPIService],
    });

    classAPIService = TestBed.inject(ClassAPIService);
    httpTestingController = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpTestingController.verify();
  });

  it('should get trending classes', () => {
    const mockedClasses: Class[] = [
      {
        upvotes: 10,
        downvotes: 2,
        upvoted: false,
        downvoted: false,
        className: 'Class 1',
        dateUpdated: new Date(),
        total_votes: 12,
      },
      {
        upvotes: 8,
        downvotes: 4,
        upvoted: false,
        downvoted: false,
        className: 'Class 2',
        dateUpdated: new Date(),
        total_votes: 12,
      },
    ];

    classAPIService.getTrendingClasses().subscribe((classes) => {
      expect(classes).to.deep.equal(mockedClasses);
    });

    const req = httpTestingController.expectOne(`${apiUrl}/getTrendingClasses`);
    expect(req.request.method).to.equal('GET');
    req.flush(mockedClasses);
  });

  it('should add a class', () => {
    const className = 'New Class';
    const mockedResponse = { message: 'Class added successfully.' };

    classAPIService.addClass(className).subscribe((response) => {
      expect(response).to.deep.equal(mockedResponse);
    });

    const req = httpTestingController.expectOne(`${apiUrl}/createClass`);
    expect(req.request.method).to.equal('POST');
    expect(req.request.body).to.deep.equal({ className });
    req.flush(mockedResponse);
  });
});
