import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router'
import { Class } from 'src/app/types';

@Component({
  selector: 'app-results',
  templateUrl: './results.component.html',
  styleUrls: ['./results.component.css']
})
export class ResultsComponent implements OnInit {

  search: string = ""
  
  results: Class[] = [{
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

  constructor(private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.search = params.get('search') ?? "";
    });
    this.results = this.rankClasses(this.results);
  }

  rankClasses(results: Class[]) {
    results.sort((a, b) => (b.upvotes - b.downvotes) - (a.upvotes - a.downvotes));
    results.forEach((results, index) => results.total_votes = results.upvotes - results.downvotes);
    return results;
  }
}


