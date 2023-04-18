import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router'
import { ClassAPIService } from 'src/app/services/class.services';
import { Class } from 'src/app/types';

@Component({
  selector: 'app-results',
  templateUrl: './results.component.html',
  styleUrls: ['./results.component.css']
})
export class ResultsComponent implements OnInit {

  search: string = ""
  results: Class[] = []

  constructor(private route: ActivatedRoute, private classAPIService: ClassAPIService) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.search = params.get('search') ?? "";
      this.results = []
      this.classAPIService.getSearchResults(this.search).subscribe(data => {
        this.results = this.rankClasses(data);
      }); 
    });
  }

  rankClasses(results: Class[]) {
    results.sort((a, b) => a.total_votes - b.total_votes);
    return results;
  }
}


