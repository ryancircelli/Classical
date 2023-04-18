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

  posts: Post[] = []

  newPost: string = "";
  errorMessage: string = "";

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.class = params.get('className') ?? "";
      this.posts = []
      this.classAPIService.getClassPosts(this.class).subscribe(data => {
        this.posts = this.rankPosts(data.map(postData => ({
          ...postData,
          timePosted: new Date(parseInt(postData.timePosted) * 1000).toLocaleString()
        })));
      }); 
    });
  }

  rankPosts(results: Post[]) {
    results.sort((a, b) => b.postVotes - a.postVotes);
    return results;
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
      resolve("");
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

  upvote(id: number) {
    this.posts[id].upvoted = true;
    this.posts[id].downvoted = false;
    // this.classAPIService.increasePostVotes(this.class)
    //   .subscribe(
    //     response => {
    //      this.posts[id].upvoted = true;
    //      this.posts[id].downvoted = false;
    //     },
    //     error => {
    // <div class="toast">
    //   <div class="alert alert-info">
    //     <div>
    //       <span>Failed to upvote.</span>
    //     </div>
    //   </div>
    // </div>
    //     }
    //   );
  }

  downvote(id : number) {
    this.posts[id].upvoted = false;
    this.posts[id].downvoted = true;
    // this.classAPIService.decreasePostVotes(this.class)
    //   .subscribe(
    //     response => {
    //       this.posts[id].upvoted = false;
    //       this.posts[id].downvoted = true;
    //     },
    //     error => {
    //       //throw toast if failed
    // <div class="toast">
    //   <div class="alert alert-info">
    //     <div>
    //       <span>Failed to downvote.</span>
    //     </div>
    //   </div>
    // </div>
    //     }
    //   );
  }
  
}
