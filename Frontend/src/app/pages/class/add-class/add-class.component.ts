import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ClassAPIService } from 'src/app/services/class.services';

@Component({
  selector: 'app-add-class',
  templateUrl: './add-class.component.html',
  styleUrls: ['./add-class.component.css']
})
export class AddClassComponent {

  constructor(public classAPIService: ClassAPIService, public router: Router ) { }

  newClass: string = "";
  errorMessage: string = '';

  submitNewClass() {
    return new Promise((resolve, reject) => {
      if (this.newClass.length < 4 || this.newClass.length > 10) {
        this.errorMessage = 'Class name must be between 4 and 10 characters long';
        reject(this.errorMessage);
      } else {
        this.classAPIService.addClass(this.newClass)
          .subscribe(
            response => {
              this.router.navigate(['/class', this.newClass]);
              resolve(response);
            },
            error => {
              console.log(error);
              if (error.text === "Class with Name = cis4930 already exists") {
                this.errorMessage = 'Class with Name = cis4930 already exists';
                reject(this.errorMessage);
              }
              this.errorMessage = error.message;
              reject(this.errorMessage);
            }
          );
      }
    });
  }
  

}
