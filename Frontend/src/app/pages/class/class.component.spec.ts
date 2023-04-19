import { of } from 'rxjs';
import { ActivatedRoute, ParamMap } from '@angular/router';

import { ClassAPIService } from 'src/app/services/class.services';
import { ClassComponent } from './class.component';
import { Post } from 'src/app/types';

function createClassAPIServiceStub(): ClassAPIService {
  return {
    apiUrl: '',
    http: null,
    getClassPosts: cy.stub(),
    createClassPost: cy.stub(),
    increasePostVotes: cy.stub(),
    decreasePostVotes: cy.stub(),
  } as unknown as ClassAPIService;
}

function createActivatedRouteStub(): ActivatedRoute {
  return {
    paramMap: {
      subscribe: cy.stub(),
    },
  } as unknown as ActivatedRoute;
}

describe('ClassComponent', () => {
  let component: ClassComponent;
  let classAPIService: ClassAPIService;
  let route: ActivatedRoute;

  beforeEach(() => {
    classAPIService = createClassAPIServiceStub();
    route = createActivatedRouteStub();
    component = new ClassComponent(route, classAPIService);
  });

  it('should create the component', () => {
    expect(component).to.exist;
  });

  it('should rank posts correctly', () => {
    const posts: Post[] = [
      { postId: 1, postContent: 'https://example.com/1', postVotes: 5, timePosted: '1618558200', upvoted: false, downvoted: false, className: 'CEN3031', postName: ''},
      { postId: 2, postContent: 'https://example.com/2', postVotes: 10, timePosted: '1618558300', upvoted: false, downvoted: false, className: 'COP4600', postName: ''},
    ];

    const rankedPosts = component.rankPosts(posts);
    expect(rankedPosts[0]).to.deep.equal(posts[1]);
    expect(rankedPosts[1]).to.deep.equal(posts[0]);
  });

  it('should fetch class posts on ngOnInit', () => {
    const className = 'Test Class';
    const paramMap: ParamMap = {
      get: () => className,
      has: () => true,
      getAll: () => [],
      keys: [],
    };
    route.paramMap.subscribe = cy.stub().callsFake((callback) => callback(paramMap));

    const classPosts: Post[] = [
      { postId: 1, postContent: 'https://example.com/1', postVotes: 10, timePosted: '1618558200', upvoted: false, downvoted: false, className: 'CEN3031', postName: ''},
      { postId: 2, postContent: 'https://example.com/2', postVotes: 5, timePosted: '1618558300', upvoted: false, downvoted: false, className: 'COP4600', postName: ''},
    ];
    classAPIService.getClassPosts = cy.stub().returns(of(classPosts));

    component.ngOnInit();

    expect(route.paramMap.subscribe).to.be.called;
    expect(classAPIService.getClassPosts).to.be.calledWith(className);
    expect(component.posts).to.deep.equal(
      classPosts.map((postData) => ({
        ...postData,
        timePosted: new Date(parseInt(postData.timePosted) * 1000).toLocaleString(),
      })),
    );
  });

  it('should set errorMessage if newPost is empty when submitting', async () => {
    component.newPost = '';
    try {
      await component.submitPost();
    } catch (errorMessage) {
      expect(errorMessage).to.equal('Please provide a link!');
      expect(component.errorMessage).to.equal('Please provide a link!');
    }
  });

  it('should set errorMessage if newPost is an invalid URL when submitting', async () => {
    component.newPost = 'invalid-url';
    try {
      await component.submitPost();
    } catch (errorMessage) {
      expect(errorMessage).to.equal('Invalid Link!\nMake sure to include http:// or https://');
      expect(component.errorMessage).to.equal('Invalid Link!\nMake sure to include http:// or https://');
    }
  });

  it('should submit a new post with a valid URL', async () => {
    component.newPost = 'https://valid-url.com';
    component.class = 'Test Class';

    classAPIService.createClassPost = cy.stub().returns(of({}));

    await component.submitPost();

    expect(classAPIService.createClassPost).to.be.calledWith(component.class, component.newPost);
    expect(component.errorMessage).to.equal('');
    expect(component.newPost).to.equal('');
  });

  it('should upvote a post', () => {
    const postId = 1;
    component.class = 'Test Class';
    component.posts = [
      { postId: 1, postContent: 'https://example.com/1', postVotes: 10, timePosted: '1618558200', upvoted: false, downvoted: false, className: 'CEN3031', postName: ''},
    ];

    classAPIService.increasePostVotes = cy.stub().returns(of({}));

    component.upvote(postId);

    expect(classAPIService.increasePostVotes).to.be.calledWith(component.class, component.posts[postId].postId);
    expect(component.posts[postId].upvoted).to.be.true;
    expect(component.posts[postId].downvoted).to.be.false;
  });

  it('should reset vote of a post', () => {
    const postId = 1;
    component.class = 'Test Class';
    component.posts = [
      { postId: 1, postContent: 'https://example.com/1', postVotes: 10, timePosted: '1618558200', upvoted: true, downvoted: false, className: 'CEN3031', postName: ''},
    ];

    classAPIService.decreasePostVotes = cy.stub().returns(of({}));

    component.resetVote(postId);

    expect(classAPIService.decreasePostVotes).to.be.calledWith(component.class, component.posts[postId].postId);
    expect(component.posts[postId].upvoted).to.be.false;
    expect(component.posts[postId].downvoted).to.be.false;
  });

  it('should downvote a post', () => {
    const postId = 1;
    component.class = 'Test Class';
    component.posts = [
      { postId: 1, postContent: 'https://example.com/1', postVotes: 10, timePosted: '1618558200', upvoted: false, downvoted: false, className: 'CEN3031', postName: ''},
    ];

    classAPIService.decreasePostVotes = cy.stub().returns(of({}));

    component.downvote(postId);

    expect(classAPIService.decreasePostVotes).to.be.calledWith(component.class, component.posts[postId].postId);
    expect(component.posts[postId].upvoted).to.be.false;
    expect(component.posts[postId].downvoted).to.be.true;
  });

  it('should load data when called', () => {
    component.class = 'Test Class';

    const classPosts: Post[] = [
      { postId: 1, postContent: 'https://example.com/1', postVotes: 10, timePosted: '1618558200', upvoted: false, downvoted: false, className: 'CEN3031', postName: ''},
      { postId: 2, postContent: 'https://example.com/2', postVotes: 5, timePosted: '1618558300', upvoted: false, downvoted: false, className: 'COP4600', postName: ''},
    ];
    classAPIService.getClassPosts = cy.stub().returns(of(classPosts));

    component.loadData();

    expect(classAPIService.getClassPosts).to.be.calledWith(component.class);
    expect(component.posts).to.deep.equal(
      classPosts.map((postData) => ({
        ...postData,
        timePosted: new Date(parseInt(postData.timePosted) * 1000).toLocaleString(),
      })),
    );
  });

});


