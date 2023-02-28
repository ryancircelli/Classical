import { Component } from '@angular/core';

type Person = {
  upvotes: number;
  downvotes: number;
  upvoted: boolean;
  downvoted: boolean;
  name: string;
  dateUpdated: Date;
};

@Component({
  selector: 'app-results',
  templateUrl: './results.component.html',
  styleUrls: ['./results.component.css']
})
export class ResultsComponent {

  constructor() {}

  results: Person[] = [{
    upvotes: 4,
    downvotes: 3,
    upvoted: false,
    downvoted: false,
    name: "CEN3031",
    dateUpdated: new Date
  },
  {
    upvotes: 2,
    downvotes: 4,
    upvoted: true,
    downvoted: false,
    name: "CIS4930",
    dateUpdated: new Date
  }]

}
