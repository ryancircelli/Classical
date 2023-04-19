import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { ClassAPIService } from './class.services';
import { Class, Post } from '../types';

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
        className: 'Class 1',
        lastUpdated: "1681836593",
        total_votes: 12,
      },
      {
        upvotes: 8,
        downvotes: 4,
        className: 'Class 2',
        lastUpdated: "1681836593",
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

  it('should get search results', () => {
    const className = 'search';
    const mockedSearchResults: Class[] = [
      { upvotes: 10, downvotes: 5, className: 'Class 1', lastUpdated: "1681836593", total_votes: 15 },
      { upvotes: 20, downvotes: 10, className: 'Class 2', lastUpdated: "1681836593", total_votes: 30 },
    ];

    classAPIService.getSearchResults(className).subscribe((results) => {
      expect(results).to.deep.equal(mockedSearchResults);
    });

    const req = httpTestingController.expectOne(`${apiUrl}/getClassesByName/${className}`);
    expect(req.request.method).to.equal('GET');
    req.flush(mockedSearchResults);
  });

  it('should get class posts', () => {
    const className = 'Class 1';
    const mockedClassPosts: Post[] = [
      { postId: 1, postContent: 'https://example.com/1', postName: 'Class 1', className: 'Class 1', upvoted: true, downvoted: true, postVotes: 7, timePosted: '1681836593' },
      { postId: 2, postContent: 'https://example.com/2', postName: 'Class 1', className: 'Class 1', upvoted: false, downvoted: false, postVotes: 15, timePosted: '1681836593' },
    ];

    classAPIService.getClassPosts(className).subscribe((posts) => {
      expect(posts).to.deep.equal(mockedClassPosts);
    });

    const req = httpTestingController.expectOne(`${apiUrl}/getPostsByClassName/${className}`);
    expect(req.request.method).to.equal('GET');
    req.flush(mockedClassPosts);
  });

  it('should create a class post', () => {
    const className = 'Class 1';
    const url = 'https://example.com/3';
    const mockedResponse = { message: 'Post created successfully.' };

    classAPIService.createClassPost(className, url).subscribe((response) => {
      expect(response).to.deep.equal(mockedResponse);
    });

    const req = httpTestingController.expectOne(`${apiUrl}/createClassPost`);
    expect(req.request.method).to.equal('POST');
    expect(req.request.body).to.deep.equal({ postClassName: className, postContent: url, postName: className });
    req.flush(mockedResponse);
  });

  it('should increase post votes', () => {
    const className = 'Class 1';
    const postID = 1;
    const mockedResponse = { message: 'Post votes increased successfully.' };

    classAPIService.increasePostVotes(className, postID).subscribe((response) => {
      expect(response).to.deep.equal(mockedResponse);
    });

    const req = httpTestingController.expectOne(`${apiUrl}/increasePostVotes`);
    expect(req.request.method).to.equal('POST');
    expect(req.request.body).to.deep.equal({ postID, postClassName: className });
    req.flush(mockedResponse);
  });

  it('should decrease post votes', () => {
    const className = 'Class 1';
    const postID = 1;
    const mockedResponse = { message: 'Post votes decreased successfully.' };

    classAPIService.decreasePostVotes(className, postID).subscribe((response) => {
      expect(response).to.deep.equal(mockedResponse);
    });

    const req = httpTestingController.expectOne(`${apiUrl}/decreasePostVotes`);
    expect(req.request.method).to.equal('POST');
    expect(req.request.body).to.deep.equal({ postID, postClassName: className });
    req.flush(mockedResponse);
  });


});
