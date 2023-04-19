import { of } from 'rxjs';
import { ActivatedRoute, ParamMap } from '@angular/router';
import { ClassAPIService } from 'src/app/services/class.services';
import { ResultsComponent } from './results.component';
import { Class } from 'src/app/types';

function createActivatedRouteStub(): ActivatedRoute {
  return {
    paramMap: {
      subscribe: cy.stub(),
    },
  } as unknown as ActivatedRoute;
}

function createClassAPIServiceStub(): ClassAPIService {
  return {
    apiUrl: '',
    http: null,
    getSearchResults: cy.stub(),
  } as unknown as ClassAPIService;
}

describe('ResultsComponent', () => {
  let component: ResultsComponent;
  let route: ActivatedRoute;
  let classAPIService: ClassAPIService;

  beforeEach(() => {
    route = createActivatedRouteStub();
    classAPIService = createClassAPIServiceStub();
    component = new ResultsComponent(route, classAPIService);
  });

  it('should create the component', () => {
    expect(component).to.exist;
  });

  it('should fetch search results on ngOnInit', () => {
    const searchResults: Class[] = [
      { upvotes: 10, downvotes: 5, className: 'Class 1', lastUpdated: "1681836593", total_votes: 15 },
      { upvotes: 20, downvotes: 10, className: 'Class 2', lastUpdated: "1681836593", total_votes: 30 },
    ];

    route.paramMap.subscribe = cy.stub().callsFake((callback: (params: ParamMap) => void) => {
      const paramMap: ParamMap = {
        has: (key: string) => key === 'search',
        get: (key: string) => key === 'search' ? 'search' : null,
        getAll: (key: string) => key === 'search' ? ['search'] : [],
        keys: ['search'],
      };
      callback(paramMap);
    });
    classAPIService.getSearchResults = cy.stub().returns(of(searchResults));

    component.ngOnInit();

    expect(route.paramMap.subscribe).to.be.called;
    expect(classAPIService.getSearchResults).to.be.calledWith('search');

    const searchResultsOutput: Class[] = [
      { upvotes: 20, downvotes: 10, className: 'Class 2', lastUpdated: "4/18/2023, 12:49:53 PM", total_votes: 30 },
      { upvotes: 10, downvotes: 5, className: 'Class 1', lastUpdated: "4/18/2023, 12:49:53 PM", total_votes: 15 },
    ];
    expect(component.results).to.deep.equal(searchResultsOutput);
  });

  it('should handle empty search results', () => {
    route.paramMap.subscribe = cy.stub().callsFake((callback: (params: ParamMap) => void) => {
      const paramMap: ParamMap = {
        has: (key: string) => key === 'search',
        get: (key: string) => key === 'search' ? 'search' : null,
        getAll: (key: string) => key === 'search' ? ['search'] : [],
        keys: ['search'],
      };
      callback(paramMap);
    });
    classAPIService.getSearchResults = cy.stub().returns(of([]));

    component.ngOnInit();

    expect(route.paramMap.subscribe).to.be.called;
    expect(classAPIService.getSearchResults).to.be.calledWith('search');
    expect(component.results).to.deep.equal([]);
  });

  it('should correctly rank classes by total_votes', () => {
    const unrankedClasses: Class[] = [
      { upvotes: 10, downvotes: 5, className: 'Class 1', lastUpdated: "1681836593", total_votes: 15 },
      { upvotes: 20, downvotes: 10, className: 'Class 2', lastUpdated: "1681836593", total_votes: 30 },
      { upvotes: 5, downvotes: 2, className: 'Class 3', lastUpdated: "1681836593", total_votes: 7 },
    ];

    const rankedClasses: Class[] = [
      { upvotes: 20, downvotes: 10, className: 'Class 2', lastUpdated: "1681836593", total_votes: 30 },
      { upvotes: 10, downvotes: 5, className: 'Class 1', lastUpdated: "1681836593", total_votes: 15 },
      { upvotes: 5, downvotes: 2, className: 'Class 3', lastUpdated: "1681836593", total_votes: 7 },
    ];

    const result = component.rankClasses(unrankedClasses);
    expect(result).to.deep.equal(rankedClasses);
  });


});
