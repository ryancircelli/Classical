import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router'

type Class = {
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
export class ResultsComponent implements OnInit {

  search: string = ""

  constructor(private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.search = params.get('search') ?? "";
    });
  }

  results: Class[] = [{
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