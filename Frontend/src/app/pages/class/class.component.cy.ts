import { of } from 'rxjs';
import { ActivatedRoute, convertToParamMap, ParamMap } from '@angular/router';

import { ClassAPIService } from '../../services/class.services';
import { ClassComponent } from './class.component';
import { Post } from '../../types';

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
  let activatedRoute: ActivatedRoute;

  beforeEach(() => {
    classAPIService = createClassAPIServiceStub();
    activatedRoute = createActivatedRouteStub();
    component = new ClassComponent(activatedRoute, classAPIService);
  });


  it('should submit a new post with a valid URL', () => {
    component.newPost = 'https://valid-url.com';
    component.class = 'Test Class';
  
    classAPIService.createClassPost = cy.stub().returns(of({}));
  
    component.submitPost().then(() => {
      expect(classAPIService.createClassPost).to.be.calledWith(component.class, component.newPost);
      expect(component.errorMessage).to.equal('');
      expect(component.newPost).to.equal('');
    });
  });
  

  it('should upvote a post', () => {
    const postId = 1;
    component.class = 'Test Class';
    component.posts = [
      { postId: 1, postContent: 'https://example.com/1', postVotes: 10, timePosted: '1618558200', upvoted: false, downvoted: false, className: 'CEN3031', postName: '' },
    ];

    classAPIService.increasePostVotes = cy.stub().returns(of({}));

    component.upvote(0); // Use the index of the post in the array

    expect(classAPIService.increasePostVotes).to.be.calledWith(component.class, component.posts[0].postId);
    expect(component.posts[0].upvoted).to.be.true;
    expect(component.posts[0].downvoted).to.be.false;
  });

  it('should reset vote of a post', () => {
    const postId = 1;
    component.class = 'Test Class';
    component.posts = [
      { postId: 1, postContent: 'https://example.com/1', postVotes: 10, timePosted: '1618558200', upvoted: true, downvoted: false, className: 'CEN3031', postName: '' },
    ];

    classAPIService.decreasePostVotes = cy.stub().returns(of({}));

    component.resetVote(0); // Use the index of the post in the array

    expect(classAPIService.decreasePostVotes).to.be.calledWith(component.class, component.posts[0].postId);
    expect(component.posts[0].upvoted).to.be.false;
    expect(component.posts[0].downvoted).to.be.false;
});

it('should create the component', () => {
    expect(component).to.exist;
  });

it('should rank posts correctly', () => {
const posts: Post[] = [
    { postId: 1, postContent: 'https://example.com/1', postVotes: 5, timePosted: '1618558200', upvoted: false, downvoted: false, className: 'CEN3031', postName: '' },
    { postId: 2, postContent: 'https://example.com/2', postVotes: 10, timePosted: '1618558300', upvoted: false, downvoted: false, className: 'COP4600', postName: ''},
];

const postsCorrect: Post[] = [
    { postId: 2, postContent: 'https://example.com/2', postVotes: 10, timePosted: '1618558300', upvoted: false, downvoted: false, className: 'COP4600', postName: ''},
    { postId: 1, postContent: 'https://example.com/1', postVotes: 5, timePosted: '1618558200', upvoted: false, downvoted: false, className: 'CEN3031', postName: '' },
];

const rankedPosts = component.rankPosts(posts);
expect(rankedPosts).to.deep.equal(postsCorrect);
});

it('should fetch class posts on ngOnInit', () => {
    const classPosts: Post[] = [
        { postId: 2, postContent: 'https://example.com/2', postVotes: 5, timePosted: '1618558300', upvoted: false, downvoted: false, className: 'COP4600', postName: ''},
        { postId: 1, postContent: 'https://example.com/1', postVotes: 10, timePosted: '1618558200', upvoted: false, downvoted: false, className: 'CEN3031', postName: '' },
    ];

activatedRoute.paramMap.subscribe = cy.stub().callsFake((callback: (params: ParamMap) => void) => {
    const paramMap: ParamMap = {
      has: (key: string) => key === 'className',
      get: (key: string) => key === 'className' ? 'className' : null,
      getAll: (key: string) => key === 'className' ? ['TclassName'] : [],
      keys: ['className'],
    };
    callback(paramMap);
  });


classAPIService.getClassPosts = cy.stub().returns(of(classPosts));

component.ngOnInit();

expect(activatedRoute.paramMap.subscribe).to.be.called;
expect(classAPIService.getClassPosts).to.be.calledWith('className');

const classPostsOutput: Post[] = [
    { postId: 1, postContent: 'https://example.com/1', postVotes: 10, timePosted: '4/16/2021, 3:30:00 AM', upvoted: false, downvoted: false, className: 'CEN3031', postName: '' },
    { postId: 2, postContent: 'https://example.com/2', postVotes: 5, timePosted: '4/16/2021, 3:31:40 AM', upvoted: false, downvoted: false, className: 'COP4600', postName: ''},
];
expect(component.posts).to.deep.equal(classPostsOutput);
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

it('should downvote a post', () => {
  const postId = 1;
  component.class = 'Test Class';
  component.posts = [
    { postId: 1, postContent: 'https://example.com/1', postVotes: 10, timePosted: '1618558200', upvoted: false, downvoted: false, className: 'CEN3031', postName: ''},
  ];

  classAPIService.decreasePostVotes = cy.stub().returns(of({}));

  component.downvote(0); // Use the index of the post in the array

  expect(classAPIService.decreasePostVotes).to.be.calledWith(component.class, component.posts[0].postId);
  expect(component.posts[0].upvoted).to.be.false;
  expect(component.posts[0].downvoted).to.be.true;
});

it('should load posts data', () => {
  component.class = 'Test Class';
  const classPosts: Post[] = [
    { postId: 1, postContent: 'https://example.com/1', postVotes: 10, timePosted: '1618558200', upvoted: false, downvoted: false, className: 'CEN3031', postName: '' },
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

