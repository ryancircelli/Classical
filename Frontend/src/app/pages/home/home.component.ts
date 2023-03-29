import { Component } from '@angular/core';

import { ClassAPIService } from '../../services/class.services'
import { Class } from '../../types'

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {

  trending: Class[] = []

  constructor(private classAPIService: ClassAPIService ) { }

  ngOnInit(): void {
    this.classAPIService.getTrendingClasses().subscribe(data => {
      this.trending = data;
      console.log(data)
    });
  }
}
