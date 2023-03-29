import { of } from 'rxjs';

import { ClassAPIService } from '../../services/class.services';
import { HomeComponent } from './home.component';
import { Class } from '../../types';

function createClassAPIServiceStub(): ClassAPIService {
  return {
    apiUrl: '',
    http: null,
    getTrendingClasses: cy.stub(),
  } as unknown as ClassAPIService;
}

describe('HomeComponent', () => {
  let component: HomeComponent;
  let classAPIService: ClassAPIService;

  beforeEach(() => {
    classAPIService = createClassAPIServiceStub();
    component = new HomeComponent(classAPIService);
  });

  it('should create the component', () => {
    expect(component).to.exist;
  });

  it('should fetch trending classes on ngOnInit', () => {
    const trendingClasses: Class[] = [
      { upvotes: 10, downvotes: 5, upvoted: false, downvoted: false, className: 'Class 1', dateUpdated: new Date(), total_votes: 15 },
      { upvotes: 20, downvotes: 10, upvoted: false, downvoted: false, className: 'Class 2', dateUpdated: new Date(), total_votes: 30 },
    ];

    classAPIService.getTrendingClasses = cy.stub().returns(of(trendingClasses));

    component.ngOnInit();

    expect(classAPIService.getTrendingClasses).to.be.called;
    expect(component.trending).to.deep.equal(trendingClasses);
  });

  it('should handle empty trending classes', () => {
    classAPIService.getTrendingClasses = cy.stub().returns(of([]));

    component.ngOnInit();

    expect(classAPIService.getTrendingClasses).to.be.called;
    expect(component.trending).to.deep.equal([]);
  });

});