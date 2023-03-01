import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { ClassAPIService } from './services/class.services'
import { Class } from './types'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Classical';
  searchTerm: string = "";
  trending: Class[] = []

  constructor(private router: Router, private classAPIService: ClassAPIService ) { }
  
  ngOnInit(): void {
    this.classAPIService.getClasses().subscribe(data => {
      this.trending = data;
      console.log(data)
    });
  }

  onSubmit() {
    if (this.searchTerm) {
      this.router.navigate(['/results', this.searchTerm]);
    }
  }

}
