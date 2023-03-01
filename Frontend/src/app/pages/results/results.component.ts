import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router'

type Class = {
  upvotes: number;
  downvotes: number;
  upvoted: boolean;
  downvoted: boolean;
  name: string;
  dateUpdated: Date;
  netVotes: number;
};

@Component({
  selector: 'app-results',
  templateUrl: './results.component.html',
  styleUrls: ['./results.component.css']
})
export class ResultsComponent implements OnInit {

  search: string = ""

  constructor(private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.search = params.get('search') ?? "";
    });
    this.rankClasses();
  }

  results: Class[] = [{
    upvotes: 2,
    downvotes: 4,
    upvoted: true,
    downvoted: false,
    name: "CIS4930",
    dateUpdated: new Date,
    netVotes: 0
  },
  {
    upvotes: 4,
    downvotes: 3,
    upvoted: false,
    downvoted: false,
    name: "CEN3031",
    dateUpdated: new Date,
    netVotes: 0
  }]
  

  rankClasses() {

    this.results.sort((a, b) => (b.upvotes - b.downvotes) - (a.upvotes - a.downvotes));
    
    this.results.forEach((results, index) => results.netVotes = results.upvotes - results.downvotes);
  }
  
  
}


