import { ResultsComponent } from "./results.component";
import { TestBed } from "@angular/core/testing";
import { Class } from 'src/app/types';
import { expect } from 'chai';
import { RouterTestingModule } from "@angular/router/testing";

describe('rankClasses()', () => {
    let resultsComponent: ResultsComponent;

    beforeEach(()=>{
      TestBed.configureTestingModule({
          imports: [RouterTestingModule],
          providers: [ResultsComponent]
        })

        resultsComponent = TestBed.inject(ResultsComponent);
    })


    it('sorts results by net votes', () => {
      // Create an example array of results with different upvotes and downvotes
      const results: Class[] = [{
        upvotes: 2,
        downvotes: 4,
        upvoted: true,
        downvoted: false,
        className: "CIS4930",
        dateUpdated: new Date,
        total_votes: 0
      },
      {
        upvotes: 4,
        downvotes: 3,
        upvoted: false,
        downvoted: false,
        className: "CEN3031",
        dateUpdated: new Date,
        total_votes: 0
      }]
  
      // Call the function to sort the results
      let output = resultsComponent.rankClasses(results);

      const expected: Class[] = [
        {
            upvotes: 4,
            downvotes: 3,
            upvoted: false,
            downvoted: false,
            className: "CEN3031",
            dateUpdated: new Date,
            total_votes: 1
        },{
            upvotes: 2,
            downvotes: 4,
            upvoted: true,
            downvoted: false,
            className: "CIS4930",
            dateUpdated: new Date,
            total_votes: -2
        }
      ]
  
      // Assert that the results are sorted by net votes in descending order
      expect(output).to.deep.eq(expected);
    });
  });
  