import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ClassAPIService } from 'src/app/services/class.services';

@Component({
  selector: 'app-add-class',
  templateUrl: './add-class.component.html',
  styleUrls: ['./add-class.component.css']
})
export class AddClassComponent {

  constructor(private classAPIService: ClassAPIService, private router: Router ) { }

  newClass: string = "";
  errorMessage: string = '';

  submitNewClass() {
    // this.classAPIService.addClass(this.newClass);
    if (this.newClass.length < 4 || this.newClass.length > 10) {
      this.errorMessage = 'Class name must be between 4 and 10 characters long';
    } else {
      this.classAPIService.addClass(this.newClass)
        .subscribe(
          response => {
            this.router.navigate(['/class', this.newClass]);
          },
          error => {
            console.log(error)
            if (error.text === "Class with Name = cis4930 already exists") {
              this.errorMessage = 'Class with Name = cis4930 already exists';
              return;
            }
            this.errorMessage = error.message;
          }
        );
    }
  }

}
