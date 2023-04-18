import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router'
import { ClassAPIService } from 'src/app/services/class.services';
import { Post } from 'src/app/types';
import { isWebUri } from 'valid-url';

@Component({
  selector: 'app-class',
  templateUrl: './class.component.html',
  styleUrls: ['./class.component.css']
})
export class ClassComponent {

  constructor(private route: ActivatedRoute, public classAPIService: ClassAPIService) { }

  class: string = ""

  posts: Post[] = [
    {
      postID: 1,
      className: "cis400",
      postVotes: 5,
      postName: "gay",
      postContent: "https://www.google.com"
    },
    {
      postID: 2,
      className: "cis400",
      postVotes: 2,
      postName: "wow",
      postContent: "https://www.bing.com"
    }
  ]

  newPost: string = "";
  errorMessage: string = "";

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.class = params.get('className') ?? "";
    });
  }

  submitPost() {
    return new Promise((resolve, reject) => {
      if (this.newPost === "") {
        this.errorMessage = 'Please provide a link!';
        this.newPost = "";
        reject(this.errorMessage);
        return;
      } else if (isWebUri(this.newPost) === undefined) {
        this.errorMessage = 'Invalid Link!\nMake sure to include http:// or https://';
        this.newPost = "";
        reject(this.errorMessage);
        return;
      }

      console.log(this.newPost, isWebUri(this.newPost))
      this.newPost = "";
      this.errorMessage = ""
      resolve("gay");
      // this.classAPIService.createPost(this.newPost)
      //   .subscribe(
      //     response => {
      //       this.router.navigate(['/class', this.newClass]);
      //       resolve(response);
      //     },
      //     error => {
      //       console.log(error);
      //       if (error.text === "Class with Name = cis4930 already exists") {
      //         this.errorMessage = 'Class with Name = cis4930 already exists';
      //         reject(this.errorMessage);
      //       }
      //       this.errorMessage = error.message;
      //       reject(this.errorMessage);
      //     }
      //   );
    });
  }

}
