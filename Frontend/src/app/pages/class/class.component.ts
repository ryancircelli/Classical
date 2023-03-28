import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router'

@Component({
  selector: 'app-class',
  templateUrl: './class.component.html',
  styleUrls: ['./class.component.css']
})
export class ClassComponent {

  constructor(private route: ActivatedRoute) { }

  class: string = ""

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.class = params.get('className') ?? "";
    });
  }

}
